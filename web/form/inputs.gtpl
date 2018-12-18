<html>
	<head>
	<title>Test</title>
	</head>
	<body>
		<form action="/inputs" method="post">
			Username:<input type="text" name="username">
            <br />
            <select name="fruit">
                <option value="apple">apple</option>
                <option value="pear">pear</option>
                <option value="banana">banana</option>
            </select>
            <br />
            <input type="radio" name="gender" value="1">Male
	        <input type="radio" name="gender" value="2">Female
            <br />
            <input type="checkbox" name="sport" value="football">Football
            <input type="checkbox" name="sport" value="basketball">Basketball
            <input type="checkbox" name="sport" value="tennis">Tennis
            <br />
			<input type="submit" value="Login">
		</form>
	</body>
</html>