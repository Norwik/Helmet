IMAGE_ID=ghcr.io/${{ github.repository }}/$IMAGE_NAME

# Change all uppercase to lowercase
IMAGE_ID=$(echo $IMAGE_ID | tr '[A-Z]' '[a-z]')

# Strip git ref prefix from version
VERSION=$(echo "${{ github.ref }}" | sed -e 's,.*/\(.*\),\1,')

# Strip "v" prefix from tag name
[[ "${{ github.ref }}" == "refs/tags/"* ]] && VERSION=$(echo $VERSION | sed -e 's/^v//')

# Use Docker `latest` tag convention
[ "$VERSION" == "main" ] && VERSION=latest

echo IMAGE_ID=$IMAGE_ID
echo VERSION=$VERSION

docker tag $IMAGE_NAME $IMAGE_ID:$VERSION
docker push $IMAGE_ID:$VERSION
