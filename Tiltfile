# Get image name and tag from environment variables or set default values
image_name = os.getenv('DOCKER_IMAGE_NAME', 'ghcr.io/ci4rail/go-template')
image_tag = os.getenv('DOCKER_IMAGE_TAG', 'latest')

full_image_name = image_name + ":" + image_tag

# Docker build with the specified image name and tag
docker_build(
    full_image_name,
    '.',
    dockerfile='Dockerfile',
    only=['api', 'cmd', 'internal', 'go.mod', 'go.sum'],
)

# Read the docker-compose configuration
docker_compose('manifest/docker-compose.yml')

# Watch specific config files and directories
watch_file('go.mod')
watch_file('go.sum')
watch_file('Dockerfile')
watch_file('manifest')

# Optionally watch the entire source code to detect changes
watch_file('api')
watch_file('cmd')
watch_file('internal')

# Configure port forwarding if needed
# port_forward('go-template', 8080, 8080)
