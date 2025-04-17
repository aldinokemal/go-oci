release:
	# Accept parameter tag
	# Example: make release tag=v1.0.0
	git tag -a $(tag) -m "$(tag)"
	# Release
	goreleaser release