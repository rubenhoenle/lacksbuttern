<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>HTMX GOLANG TEST</title>
  <script src="https://unpkg.com/htmx.org@1.9.8" integrity="sha384-rgjA7mptc2ETQqXoYC3/zJvkU7K/aP44Y+z7xQuJiVnB/422P/Ak+F/AqFR7E4Wr" crossorigin="anonymous"></script>
</head>
<body>
<h1>{{ .timestamp }}</h1>
<div hx-get="/lazy" hx-trigger="load">
</div>
</body>
</html>
