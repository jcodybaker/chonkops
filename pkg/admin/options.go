package admin

import "io/fs"

type Option func(*Server)

// WithAssetFileSystem sets a file-system (real or virtual) for serving static assets.
func WithAssetFileSystem(assetFS fs.FS) Option {
	return func(s *Server) {
		s.assetFS = assetFS
	}
}
