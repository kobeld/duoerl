$(function() {
	// Upload images
	$('#fileupload').fileupload({
        dataType: 'json',
        done: function (e, data) {
			var self = this;
			if(data.result.Attachments == null || data.result.Attachments.length == 0) {
				return;
			}
			_.each(data.result.Attachments, function (att) {
				var imgUrl = "/img/" + att.Id + "/" + att.Filename;
				$(".object-img").attr("src", imgUrl);
				$(".upload-img-input").val(imgUrl);
			});
		}
	});

});