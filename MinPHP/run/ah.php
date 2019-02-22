<?php  
$url = $_POST["url"];
$parms = $_POST["dataparam"];
$posttype = $_POST["posttype"];

$postdate = http_build_query($parms);

$opts = array (
		'http' => array (
			'method' => "POST",
			'header'=> "Content-type: application/x-www-form-urlencoded",
			'content' => $postdate
		)
	);

if($posttype =="GET"){
	$url = $url."?" . $postdate;
	$opts = array (
		'http' => array (
			'method' => 'GET'
		)
	);
}

$context = stream_context_create($opts);
$html = file_get_contents($url , false, $context);

echo $html;
//echo json_encode($html ."/n". $url ."/n". $context."/n". json_encode($parms);  
?> 