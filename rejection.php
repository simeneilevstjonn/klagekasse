To: &lt;<?php echo $argv[2]; ?>&gt;
From: &lt;noreply@trok.no&gt;
Subject: <?php echo $argv[3]; ?>
Content-type: text/html

<!doctype HTML>
<html>
<head>
    <meta charset="UTF-8">
</head>
<body>
<p>
    Dear customer
</p>

<p>
    We are referring to your inquiry with ID <strong><?php echo $argv[1]; ?></strong>.
</p>

<p>
    We have evaluated your inquiry, and found that it is not applicable.
</p>

<p>
    Your inquiry has been closed.<br>
    Please open a new inquiry if you still need support.
</p>

<p>
    Best regards,<br>
    <br>
    Company Name
</p>
</body>
</html>