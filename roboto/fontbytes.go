package roboto

import _ "embed"

//go:embed Roboto-Black.ttf
var fontRobotoBlack []byte

func RobotoBlack() []byte { return fontRobotoBlack }

//go:embed Roboto-BlackItalic.ttf
var fontRobotoBlackItalic []byte

func RobotoBlackItalic() []byte { return fontRobotoBlackItalic }

//go:embed Roboto-Bold.ttf
var fontRobotoBold []byte

func RobotoBold() []byte { return fontRobotoBold }

//go:embed Roboto-BoldItalic.ttf
var fontRobotoBoldItalic []byte

func RobotoBoldItalic() []byte { return fontRobotoBoldItalic }

//go:embed Roboto-Italic.ttf
var fontRobotoItalic []byte

func RobotoItalic() []byte { return fontRobotoItalic }

//go:embed Roboto-Light.ttf
var fontRobotoLight []byte

func RobotoLight() []byte { return fontRobotoLight }

//go:embed Roboto-LightItalic.ttf
var fontRobotoLightItalic []byte

func RobotoLightItalic() []byte { return fontRobotoLightItalic }

//go:embed Roboto-Medium.ttf
var fontRobotoMedium []byte

func RobotoMedium() []byte { return fontRobotoMedium }

//go:embed Roboto-MediumItalic.ttf
var fontRobotoMediumItalic []byte

func RobotoMediumItalic() []byte { return fontRobotoMediumItalic }

//go:embed Roboto-Regular.ttf
var fontRobotoRegular []byte

func RobotoRegular() []byte { return fontRobotoRegular }

//go:embed Roboto-Thin.ttf
var fontRobotoThin []byte

func RobotoThin() []byte { return fontRobotoThin }

//go:embed Roboto-ThinItalic.ttf
var fontRobotoThinItalic []byte

func RobotoThinItalic() []byte { return fontRobotoThinItalic }
