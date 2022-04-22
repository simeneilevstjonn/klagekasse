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
            Thank you for your inquiry.
        </p>

        <p>
            Your case has been assigned case ID <strong><?php echo $argv[1]; ?></strong>.<br>
            Please refer to this ID for all further communication.<br>
            We will get back to you within five business days.
        </p>

        <p>
            Best regards,<br>
            <br>
            Company Name
        </p>
    </body>
</html>