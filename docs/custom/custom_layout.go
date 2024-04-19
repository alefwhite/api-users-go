package custom

import "fmt"

var JS = fmt.Sprintf(`
   // set custom title
    document.title = 'Swagger Dark Mode With Go';

    // dark mode
    const style = document.createElement('style');
    style.innerHTML = %s;
    document.head.appendChild(style);
  `, "`"+customCSS+"`")
