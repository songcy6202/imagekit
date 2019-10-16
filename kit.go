package imagekit

import(
	"github.com/docker/distribution"
)


// ImageKit represents toolset for manipulate docker image
type ImageKit interface {
	Packer
	StaticBuilder
	ImageDescriptor
}

// Packer pack layers to generate a legal docker image
type Packer interface {
	Pack()
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