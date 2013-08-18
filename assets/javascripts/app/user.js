$(function() {

	var birthday = $("#birthday-picker").data("birthday");

	$("#birthday-picker").birthdaypicker({
		dateFormat: "bigEndian",
		monthFormat: "number",
		maxAge: 80,
		fieldName: "Profile.Birthday",
		defaultDate: birthday,
	});
});