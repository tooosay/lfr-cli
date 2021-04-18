package project

import "encoding/xml"

const (
	XMLHeader = `<?xml version="1.0"?>` + "\n"
)

type Pom struct {
	XMLName        xml.Name `xml:"project"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xmlns:schemaLocation,attr"`
	ModelVersion   string   `xml:"modelVersion"`
	Parent         struct {
		GroupId      string `xml:"groupId"`
		ArtifactId   string `xml:"artifactId"`
		Version      string `xml:"version"`
		RelativePath string `xml:"relativePath"`
	} `xml:"parent"`
	ArtifactId string `xml:"artifactId"`
	Name       string `xml:"name"`
	Packaging  string `xml:"packaging"`
	Modules    struct {
		Module []string `xml:"module"`
	} `xml:"modules"`
}

type WorkspacePom struct {
	XMLName        xml.Name `xml:"project"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xml:"xmlns:schemaLocation,attr"`
	ModelVersion   string   `xml:"modelVersion"`
	Parent         struct {
		GroupId      string `xml:"groupId"`
		ArtifactId   string `xml:"artifactId"`
		Version      string `xml:"version"`
		RelativePath string `xml:"relativePath"`
	} `xml:"parent"`
	ArtifactId string `xml:"artifactId"`
	Name       string `xml:"name"`
	Packaging  string `xml:"packaging"`
	Modules    struct {
		Module []string `xml:"module"`
	} `xml:"modules"`
	Properties struct {
		LiferayBomVersion          string `xml:"liferay.bom.version"`
		LiferayDockerImage         string `xml:"liferay.docker.image"`
		LiferayWorkspaceBundleURL  string `xml:"liferay.workspace.bundle.url"`
		LiferayRepositoryURL       string `xml:"liferay.repository.url"`
		ProjectBuildSourceEncoding string `xml:"project.build.sourceEncoding"`
	} `xml:"properties"`
}