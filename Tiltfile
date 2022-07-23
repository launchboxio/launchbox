docker_build("launchboxio/server:latest", ".")

k8s_yaml(helm('deploy/charts/launchbox',
  name='launchbox',
  namespace='launchbox-system',
  values=['./deploy/dev.values.yaml'],
  set=[
    'image.repository=launchboxio/server',
    'image.tag=latest'
  ]
))

watch_file('deploy/charts/launchbox')

allow_k8s_contexts('default')