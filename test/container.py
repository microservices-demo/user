import argparse
import sys
import unittest
import requests

from util.Api import Api
from util.Dredd import Dredd
from util.Docker import Docker
from time import sleep

class ContainerTest(unittest.TestCase):
    container_name = '{0}-dev'.format(SERVICE)
    TAG = "latest"
    def __init__(self, methodName='runTest'):

    def setUp(self):
        self.ip = Docker().get_container_ip(UserContainerTest.container_name)

    def test_api_validated(self):
        self.wait_or_fail('http://'+ self.ip +':8084/{0}'.format(SERVICEUP))
        out = Dredd().test_against_endpoint(SERVICE, self.container_name, "http://{0}/".format(SERVICE))
        self.assertGreater(out.find("0 failing"), -1)
        self.assertGreater(out.find("0 errors"), -1)
        print(out)

    def wait_or_fail(self,endpoint, limit=20):
        while Api().noResponse(endpoint):
            if limit == 0:
                self.fail("Couldn't get the API running")
                limit = limit - 1
                sleep(1)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--tag', default="latest", help='The tag of the image to use. (default: latest)')
    parser.add_argument('--service', default="", help='The service name')
    parser.add_argument('--serviceup', default="", help='The service up endpoint')
    parser.add_argument('unittest_args', nargs='*')
    args = parser.parse_args()
    ContainerTest.TAG = args.tag
    ContainerTest.SERVICE = args.service
    ContainerTest.SERVICEUP = args.serviceup
    # Now set the sys.argv to the unittest_args (leaving sys.argv[0] alone)
    sys.argv[1:] = args.unittest_args
    unittest.main()
