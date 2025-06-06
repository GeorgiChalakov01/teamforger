// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.887
package sections

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SignUpForm() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container\"><div class=\"row\"><h1 class=\"col\">SignUp</h1></div><div class=\"row\"><form class=\"row g-3 needs-validation\" action=\"/signup-validation\" method=\"post\" novalidate><!-- Email Field --><div class=\"col-12\"><label for=\"email\" class=\"form-label\">Employee Email address</label> <input type=\"email\" class=\"form-control\" name=\"email\" id=\"email\" placeholder=\"Enter email\" required><div class=\"invalid-feedback\">Please provide a valid email address.</div></div><!-- Password Field --><div class=\"col-12\"><label for=\"password\" class=\"form-label\">Password</label> <input type=\"password\" class=\"form-control\" name=\"password\" id=\"password\" placeholder=\"Password\" required><div class=\"invalid-feedback\">Please enter a password.</div></div><!-- Repeated Password Field --><div class=\"col-12\"><label for=\"repeatedPassword\" class=\"form-label\">Repeat Password</label> <input type=\"password\" class=\"form-control\" name=\"repeatedPassword\" id=\"repeatedPassword\" placeholder=\"Repeat Password\" required><div class=\"invalid-feedback\" id=\"passwordMatchFeedback\">Passwords must match.</div></div><!-- Submit Button --><div class=\"col-12\"><button type=\"submit\" class=\"btn btn-primary\">Sign Up</button></div></form></div></div><!-- Validation Script --><script>\n\t(() => {\n\t\t'use strict'\n\n\t\t// Password matching validation\n\t\tconst password = document.getElementById('password')\n\t\tconst repeatedPassword = document.getElementById('repeatedPassword')\n\t\tconst passwordFeedback = document.getElementById('passwordMatchFeedback')\n\t\t\n\t\tfunction validatePassword() {\n\t\t\tif (repeatedPassword.value !== password.value) {\n\t\t\t\trepeatedPassword.setCustomValidity('Passwords do not match')\n\t\t\t\tpasswordFeedback.textContent = 'Passwords must match.'\n\t\t\t} else {\n\t\t\t\trepeatedPassword.setCustomValidity('')\n\t\t\t\tpasswordFeedback.textContent = 'Please repeat your password.'\n\t\t\t}\n\t\t}\n\n\t\tpassword.addEventListener('input', validatePassword)\n\t\trepeatedPassword.addEventListener('input', validatePassword)\n\n\t\t// Bootstrap validation\n\t\tconst forms = document.querySelectorAll('.needs-validation')\n\n\t\tArray.from(forms).forEach(form => {\n\t\t\tform.addEventListener('submit', event => {\n\t\t\t\tvalidatePassword() // Check passwords before submission\n\t\t\t\t\n\t\t\t\tif (!form.checkValidity()) {\n\t\t\t\t\tevent.preventDefault()\n\t\t\t\t\tevent.stopPropagation()\n\t\t\t\t}\n\n\t\t\t\tform.classList.add('was-validated')\n\t\t\t}, false)\n\t\t})\n\t})()\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
