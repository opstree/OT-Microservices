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
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css">
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"></script>
    </head>
    <body>
        <h1>OpsTree Golang Mysql Curd Example</h1>
    <div class="form-row">
{{ end }}

{{ define "Footer" }}
</div>
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
   <form>
   <div class="form-group">
      <label for="inputName">Name</label>
      <input type="text" class="form-control" id="inputName" placeholder="Name" name="name">
    </div>
  <div class="form-group">
    <label for="inputCity">City</label>
    <input type="text" class="form-control" id="inputCity" placeholder="Name" name="city">
  </div>
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
