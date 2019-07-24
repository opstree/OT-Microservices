package webapp

const htmltemplate=`{{ define "Index" }}
{{ template "Header" }}
  {{ template "Menu"  }}
  <h2> Registered </h2>
  <table border="1">
	<thead>
	<tr>
	  <td>ID</td>
	  <td>Name</td>
	  <td>City</td>
	  <td>View</td>
	  <td>Edit</td>
	  <td>Delete</td>
	</tr>
	 </thead>
	 <tbody>
  {{ range . }}
	<tr>
	  <td>{{ .Id }}</td>
	  <td> {{ .Name }} </td>
	  <td>{{ .City }} </td> 
	  <td><a href="/show?id={{ .Id }}">View</a></td>
	  <td><a href="/edit?id={{ .Id }}">Edit</a></td>
	  <td><a href="/delete?id={{ .Id }}">Delete</a><td>
	</tr>
  {{ end }}
	 </tbody>
  </table>
{{ template "Footer" }}
{{ end }}

{{ define "Header" }}
<!DOCTYPE html>
<html lang="en-US">
    <head>
        <title>OpsTree Golang Curd Example</title>
        <meta charset="UTF-8" />
    </head>
    <body>
        <h1>OpsTree Golang Mysql Curd Example</h1>   
{{ end }}

{{ define "Footer" }}
    </body>
</html>
{{ end }}

{{ define "Menu" }}
<a href="/">HOME</a> | 
<a href="/new">NEW</a>
{{ end }}

{{ define "Show" }}
  {{ template "Header" }}
    {{ template "Menu"  }}
    <h2> Register {{ .Id }} </h2>
      <p>Name: {{ .Name }}</p>
      <p>City:  {{ .City }}</p><br /> <a href="/edit?id={{ .Id }}">Edit</a></p>
  {{ template "Footer" }}
{{ end }}

{{ define "New" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2>New Name and City</h2>  
    <form method="POST" action="insert">
      <label> Name </label><input type="text" name="name" /><br />
      <label> City </label><input type="text" name="city" /><br />
      <input type="submit" value="Save user" />
    </form>
  {{ template "Footer" }}
{{ end }}

{{ define "Edit" }}
  {{ template "Header" }}
    {{ template "Menu" }} 
   <h2>Edit Name and City</h2>  
    <form method="POST" action="update">
      <input type="hidden" name="uid" value="{{ .Id }}" />
      <label> Name </label><input type="text" name="name" value="{{ .Name }}"  /><br />
      <label> City </label><input type="text" name="city" value="{{ .City }}"  /><br />
      <input type="submit" value="Save user" />
    </form><br />    
  {{ template "Footer" }}
{{ end }}`
