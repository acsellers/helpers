//
package assets

import (
	"html/template"
)

var (
	Functions = map[string]interface{}{
		"audio_path":             AudioPath,
		"audio_tag":              AudioTag,
		"favicon_link_tag":       FaviconLinkTag,
		"font_path":              FontPath,
		"image_path":             ImagePath,
		"image_tag":              ImageTag,
		"javascript_path":        JavascriptPath,
		"javascript_src_tag":     JavascriptSrcTag,
		"javascript_include_tag": JavascriptIncludeTag,
		"stylesheet_link_tag":    StylesheetLinkTag,
		"stylesheet_path":        StylesheetPath,
		"video_path":             VideoPath,
		"video_tag":              VideoTag,
	}

	Paths = map[string]string{
		"javascript": "/public/js/",
		"css":        "/public/css",
		"image":      "/public/img",
		"audio":      "/public/audio/",
		"font":       "/public/fonts/",
		"video":      "/public/video/",
	}
	PrefixedPaths = map[string]map[string]string{
		"javascript": map[string]string{
			"asset":  "/assets/js",
			"google": "//ajax.googleapis.com/ajax/libs/",
			"ms":     "http://ajax.aspnetcdn.com/ajax/",
			"cdnjs":  "//cdnjs.cloudflare.com/ajax/libs/",
		},
		"css": map[string]string{
			"cdnjs": "//cdnjs.cloudflare.com/ajax/libs/",
		},
		"image": map[string]string{
			"hold":    "http://placehold.it/",
			"lorem":   "http://lorempixel.com/",
			"kittens": "http://placekitten.com/",
			"dogs":    "http://placedog.com/",
		},
	}
)
