package libyaml

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type StartCtx struct {
	err    error
	config *RootConfig
}

func (ctx *StartCtx) buildFunctions() template.FuncMap {
	staticMap := template.FuncMap{
		"Now":          ctx.noParameterValidation,
		"NowFmt":       ctx.oneParameterValidation,
		"ToLower":      ctx.oneParameterValidation,
		"ToUpper":      ctx.oneParameterValidation,
		"UrlEncode":    ctx.oneParameterValidation,
		"Base64Encode": ctx.oneParameterValidation,
		"Base64Decode": ctx.oneParameterValidation,
		"Split":        strings.Split,
		"Add":          ctx.numericValidation,
		"Sub":          ctx.numericValidation,
		"Mult":         ctx.numericValidation,
		"Div":          ctx.numericValidation,
	}

	preStartFuncMap := template.FuncMap{
		"ConfigOption":            ctx.oneParameterValidation,
		"ConfigOptionIndex":       ctx.twoParameterValidationInt64,
		"ConfigOptionData":        ctx.oneParameterValidation,
		"ConfigOptionEquals":      ctx.twoParameterValidationRetBool,
		"ConfigOptionNotEquals":   ctx.twoParameterValidationRetBool,
		"LicenseFieldValue":       ctx.oneParameterValidation,
		"LicenseProperty":         ctx.oneParameterValidation,
		"LdapCopyAuthFrom":        ctx.oneParameterValidation,
		"ConsoleSetting":          ctx.oneParameterValidation,
		"ConsoleSettingEquals":    ctx.twoParameterValidationRetBool,
		"ConsoleSettingNotEquals": ctx.twoParameterValidationRetBool,
		"AppSetting":              ctx.oneParameterValidation,
	}

	startFuncMap := template.FuncMap{
		"NodePublicIPAddressAll":    ctx.validateComponentAndContainer,
		"NodePublicIPAddressFirst":  ctx.validateComponentAndContainer,
		"NodePublicIPAddress":       ctx.validateComponentAndContainer,
		"NodePrivateIPAddressAll":   ctx.validateComponentAndContainer,
		"NodePrivateIPAddressFirst": ctx.validateComponentAndContainer,
		"NodePrivateIPAddress":      ctx.validateComponentAndContainer,
		"ContainerExposedPortAll":   ctx.validateComponentContainerAndPort,
		"ContainerExposedPortFirst": ctx.validateComponentContainerAndPort,
		"ContainerExposedPort":      ctx.validateComponentContainerAndPort,

		// deprecated but still supported
		"HostPublicIpAddress":     ctx.validateComponentAndContainer,
		"HostPublicIpAddressAll":  ctx.validateComponentAndContainer,
		"HostPrivateIpAddress":    ctx.validateComponentAndContainer,
		"HostPrivateIpAddressAll": ctx.validateComponentAndContainer,
	}

	hostFunMap := template.FuncMap{
		"ThisNodePublicIPAddress":  ctx.noParamValidation,
		"ThisNodePrivateIPAddress": ctx.noParamValidation,
		"ThisNodeDockerAddress":    ctx.noParamValidation,

		// deprecated but still supported
		"ThisHostPublicIpAddress":  ctx.noParamValidation,
		"ThisHostPrivateIpAddress": ctx.noParamValidation,
		"ThisNodeInterfaceAddress": ctx.oneParameterValidation,
		"ThisHostInterfaceAddress": ctx.oneParameterValidation,
		"InterfaceAddress":         ctx.oneParameterValidation,
	}

	return merge(preStartFuncMap, staticMap, startFuncMap, hostFunMap)
}

func merge(maps ...map[string]interface{}) map[string]interface{} {
	retMap := make(map[string]interface{})
	for _, aMap := range maps {
		for key, val := range aMap {
			retMap[key] = val
		}
	}
	return retMap
}

func (ctx *StartCtx) noParameterValidation() string {
	return "SUCCESS"
}

func (ctx *StartCtx) oneParameterValidation(param string) string {
	return "SUCCESS"
}

func (ctx *StartCtx) twoParameterValidationInt64(param string, val int64) string {
	return "SUCCESS"
}

func (ctx *StartCtx) twoParameterValidationRetBool(one string, two string) bool {
	return true
}

func (ctx *StartCtx) noParamValidation() string {
	return "SUCCESS"
}

func (ctx *StartCtx) numericValidation(a, b interface{}) interface{} {
	return 5
}

func (ctx *StartCtx) validateComponentAndContainer(componentName string, containerName string) string {
	component := ctx.findComponent(componentName)
	if component == nil {
		ctx.err = fmt.Errorf("No such component: " + componentName)
		return ""
	}

	container := ctx.findContainer(component, containerName)
	if container == nil {
		ctx.err = fmt.Errorf("No such container: " + containerName)
		return ""
	}
	return "SUCCESS"
}

func (ctx *StartCtx) validateComponentContainerAndPort(componentName string, containerName string, portNumber string) string {
	component := ctx.findComponent(componentName)
	if component == nil {
		ctx.err = fmt.Errorf("No such component: " + componentName)
		return ""
	}

	container := ctx.findContainer(component, containerName)
	if container == nil {
		ctx.err = fmt.Errorf("No such container: " + containerName)
		return ""
	}

	// Make sure the port is a number, the port may not be explictly mapped by the target container via
	// the Replicated yaml as it can be exposed directly in the Dockerfile
	num, err := strconv.Atoi(portNumber)
	if err != nil {
		ctx.err = fmt.Errorf("Bad port number string")
		return ""
	}
	if num < 0 || num >= 65535 {
		ctx.err = fmt.Errorf("Illegal port number, must be 0 < port >= 65535")
		return ""
	}

	return "SUCCESS"
}

func (ctx *StartCtx) findComponent(componentName string) *Component {
	for _, component := range ctx.config.Components {
		if component.Name == componentName {
			return component
		}
	}
	return nil
}

func (ctx *StartCtx) findContainer(component *Component, containerName string) *Container {
	for _, container := range component.Containers {
		if container.ImageName == containerName {
			return container
		}
	}
	return nil
}
