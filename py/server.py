import logging
from concurrent.futures import ThreadPoolExecutor

import grpc

import calc_pb2 as calc__pb2
import calc_pb2_grpc as pb2__grpc


class SolveServer(pb2__grpc.SolverServicer):
    def Solve(self, request, context):
        a = request.a
        b = request.b
        sign = request.sign
        print(f"data from request: {a=}, {b=}, {sign=}")
        resp = -1
        m = {
            "+": lambda a, b: a + b,
            "-": lambda a, b: a - b,
            "*": lambda a, b: a * b,
            "/": lambda a, b: a / b,
        }
        try:
            resp = calc__pb2.SolverResponse(result=m[sign](a, b))
        except Exception as e:
            print(e)
            resp = calc__pb2.SolverResponse(result=-1)
        return resp


if __name__ == "__main__":
    server = grpc.server(ThreadPoolExecutor())
    pb2__grpc.add_SolverServicer_to_server(SolveServer(), server)
    port = 9999
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print("server start")
    server.wait_for_termination()