$(function() {

	var birthday = $("#birthday-picker").data("birthday");

	$("#birthday-picker").birthdaypicker({
		dateFormat: "bigEndian",
		monthFormat: "number",
		maxAge: 80,
		fieldName: "Profile.Birthday",
		defaultDate: birthday,
	});

	$(".post-btn").click(function(){
		var self = this,
			postId = $(this).data("post-id"),
			content = $(".post-content").val();

		$.post("/post/create", {Id: postId, Content: content})
		.done(function(data){
			console.log(data);
		}).fail(function(err){
			alert(err);
		});
	});
});