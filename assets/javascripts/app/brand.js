$(function() {

	// Add a wish item
	$(".btn-follow-brand").click(function() {
		var self = this;
		var brandId = $(this).closest(".show-brand").data("brand");
		$.post("/brand/follow", {bid: brandId})
		.done(function(){
			$(self).hide();
			$(".btn-unfollow-brand").show();
		})
		.fail(function(){ alert("Internal Error!") });
	});

	$(".btn-unfollow-brand").click(function() {
		var self = this;
		var brandId = $(this).closest(".show-brand").data("brand");
		$.post("/brand/unfollow", {bid: brandId})
		.done(function(){
			$(self).hide();
			$(".btn-follow-brand").show();
		})
		.fail(function(){ alert("Internal Error!") });
	});
});