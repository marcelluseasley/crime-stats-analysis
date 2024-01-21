import grpc
from concurrent import futures
import logging
import requests

import proto.crime_stats_pb2 as crimestats_pb2
import proto.crime_stats_pb2_grpc as crimestats_pb2_grpc
from grpc_reflection.v1alpha import reflection

import urllib.parse

class CrimeStatsServiceServicer(crimestats_pb2_grpc.CrimeStatsServiceServicer):
    def CrimeStats(self, request, context):
        logging.info('CrimeStats function was invoked with %s', request)

        address = request.street + ' ' + request.city + ' ' + request.state + ' ' + request.zipcode
        address = urllib.parse.quote_plus(address)
        url = f"http://dev.virtualearth.net/REST/v1/Locations?q={address}&key=xxx"
        response = requests.get(url)

        if response.status_code != 200:
            context.abort(grpc.StatusCode.NOT_FOUND, "Could not find location")

        location = response.json()

        if not location or not location['resourceSets'] or not location['resourceSets'][0]['resources']:
            context.abort(grpc.StatusCode.INTERNAL, "Wrong number of coordinates")

        main_resource = location['resourceSets'][0]['resources'][0]
        address = main_resource['name']
        coordinates = main_resource['point']['coordinates']
        logging.info(coordinates)
        logging.info(dir(crimestats_pb2.CrimeStatsResponse))
        return crimestats_pb2.CrimeStatsResponse(
            items=[
                crimestats_pb2.CrimeStatsResponse.Crime(
                    location=coordinates,
                    description=address,
                )
            ]
        )

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    crimestats_pb2_grpc.add_CrimeStatsServiceServicer_to_server(CrimeStatsServiceServicer(), server)

    SERVICE_NAMES = (
        crimestats_pb2.DESCRIPTOR.services_by_name['CrimeStatsService'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port('[::]:50051')
    server.start()
    logging.info("Server started. Listening on port 50051.")
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    serve()