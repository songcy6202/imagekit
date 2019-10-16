package imagekit

import(
	"github.com/docker/distribution"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/image"
)


// ImageKit represents toolset for manipulate docker image
type ImageKit interface {
	Packer
	StaticBuilder
	ImageDescriptor
}

// Packer pack layers to generate a legal docker image
type Packer interface {
	Pack(name string, tag string, imageMeta image.Image, layers []distribution.Descriptor) (distribution.Manifest, error)
}

// DockerImage Descripter
type DockerImage struct{}

// StaticBuilder build image with static files
type StaticBuilder interface {
	Build() DockerImage
}

// ImageDescriptor is a descripter for a docker image
type ImageDescriptor interface {
	//WIP
	Manifest() distribution.Manifest
}