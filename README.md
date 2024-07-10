# Template Hello
Is a small rudimentairy application that uses practical template nesting.
## Template nesting
File render/render.go provides a way GO template nesting is achieved by this snippet:

		if fName == "home.page.tmpl" {
			tmpSet, err = tmpSet.ParseGlob("./templates/*.partial.tmpl")
			if err != nil {
				fmt.Println("95 - ParseGlob error")
				return tmpCache, err
			}
		}

Maybe (probably) this is not the way best to insert nested GO templates.
But, it works and for this example it's OK.

Figuring out to let ParseGlob() do the job seems to be the next step.

Have Fun!
