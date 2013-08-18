$(function() {

	var birthday = $("#birthday-picker").data("birthday");

	$("#birthday-picker").birthdaypicker({
		dateFormat: "bigEndian",
		monthFormat: "number",
		maxAge: 80,
		fieldName: "Profile.Birthday",
		defaultDate: birthday,
	});

	// Upload avatar

	$('#fileupload').fileupload({
        dataType: 'json',
        done: function (e, data) {
			var self = this;
			if(data.result.Attachments == null || data.result.Attachments.length == 0) {
				return;
			}
			_.each(data.result.Attachments, function (att) {
				var avatarUrl = "/img/" + att.Id + "/" + att.Filename;
				$(".user-avatar").attr("src", avatarUrl);
				$(".upload-avatar-input").val(avatarUrl);
			});
		}
	});

});