<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Gin Template</title>
</head>
<body>
    <h1>Embedded HTML in Script</h1>
    <script type="text/javascript">
        const content = `
        {{.UnescapedHTML}}
        `;
        console.log(content);
        const conal = content;
        console.log(content);
    </script>
</body>
</html>