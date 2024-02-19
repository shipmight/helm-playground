version_settings(constraint='>=0.22.2')

docker_build(
    'helm-playground',
    context='.',
    dockerfile='./deploy/Dockerfile',
    live_update=[
        sync('./app/', '/usr/share/nginx/html/'),
        run('echo $?')
    ]
)

# if k8s_namespace() == 'default':
#     fail("Failing early to avoid deploying to default namespace")

k8s_yaml('deploy/helm-playground.yaml')

k8s_resource(
    'helm-playground',
    port_forwards='5000:5000',
    labels=['app']
)

k8s_resource(new_name='namespace', labels=['k8s_general'], objects=['helm-playground:namespace'])
