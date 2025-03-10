// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package contacts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/adamwoolhether/hypermedia/business/contacts/archiver"
)

func Archive(arch archiver.ArchiveView) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"archive-ui\" class=\"download-ui\" hx-target=\"this\" hx-swap=\"outerHTML\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch arch.Status {
		case archiver.Waiting:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"/contacts/archive\">Download Contact Archive</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case archiver.Running:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div style=\"position: absolute; color: black;\">Creating archive... ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Percent)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/hypermedia/web/view/contacts/archive.templ`, Line: 19, Col: 93}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("%</div><div hx-get=\"/contacts/archive\" hx-trigger=\"load delay:500ms\"><div class=\"progress\"><div id=\"archive-progress\" class=\"progress-bar\" role=\"progressbar\" aria-valuenow=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Percent)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/hypermedia/web/view/contacts/archive.templ`, Line: 25, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-progress=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(arch.Percent)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/hypermedia/web/view/contacts/archive.templ`, Line: 26, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = updateProgressBar(arch.Percent).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case archiver.Complete:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("   <a hx-boost=\"false\" href=\"/contacts/archive/file\" _=\"on load click() me\">Archive Ready! Click here to download &downarrow;</a> <button hx-delete=\"/contacts/archive\">Clear Download</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// updateProgressBar allows us to update the progress bar.
// Templ doesn't allow us to dynamically set `style` attributes,
// so we need to rely on a little bit of JS here.
func updateProgressBar(percentage string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_updateProgressBar_7253`,
		Function: `function __templ_updateProgressBar_7253(percentage){console.log(percentage);
    const progressBar = document.getElementById('archive-progress');
    progressBar.style.setProperty('width', ` + "`" + `${percentage}%` + "`" + `);
}`,
		Call:       templ.SafeScript(`__templ_updateProgressBar_7253`, percentage),
		CallInline: templ.SafeScriptInline(`__templ_updateProgressBar_7253`, percentage),
	}
}

var _ = templruntime.GeneratedTemplate
