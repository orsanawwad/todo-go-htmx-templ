// Code generated by templ@v0.2.316 DO NOT EDIT.

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if err != nil {
			return err
		}
		var_2 := `Todos`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><script src=\"/assets/js/htmx.min.js\">")
		if err != nil {
			return err
		}
		var_3 := ``
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"/assets/js/ws.js\">")
		if err != nil {
			return err
		}
		var_4 := ``
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><script src=\"/assets/js/tailwindcss.min.js\">")
		if err != nil {
			return err
		}
		var_5 := ``
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script></head><body><div class=\"relative flex min-h-screen flex-col justify-center overflow-hidden bg-blue-100 py-6 sm:py-12\"><div class=\"relative bg-white px-6 pt-10 pb-8 shadow-xl ring-1 ring-gray-900/5 sm:mx-auto sm:max-w-lg sm:rounded-lg sm:px-10\"><div hx-get=\"/todos\" hx-trigger=\"load\" hx-swap=\"outerHTML\"></div></div></div></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

// <div hx-ext="ws" ws-connect="/ws/connect">

// <div>
// 	<div>
// 		<div id="notifications"></div>
// 		<div id="chat_room">
// 		</div>
// 		<form id="form" ws-send>
// 			<input name="chat_message" />
// 		</form>
// 	</div>
// </div>
