from util.Docker import Docker


class Dredd:
    image = 'weaveworksdemos/openapi'
    container_name = ''

    def test_against_endpoint(self, service, service_container, api_endpoint):
        self.container_name = Docker().random_container_name('openapi')
        command = ['docker', 'run',
                   '-h', 'openapi',
                   '--name', self.container_name,
                   '--link', service_container,
                   '-v', "{0}:{1}".format("apispec/{0}/".format(service), "/tmp/specs/"),
                   Dredd.image,
                   "/tmp/specs/{0}.json".format(service),
                   api_endpoint,
                   "-f",
                   "/tmp/specs/hooks.js"
                   ]
        out = Docker().execute(command)
        Docker().kill_and_remove(self.container_name)
        return out
