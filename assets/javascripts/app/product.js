$(function() {

	// Add a wish item
	$(".btn-add-wishitem").click(function() {
		var self = this;
		var productId = $(this).closest(".show-product").data("product");
		$.post("/wish_item/add", {pid: productId})
		.done(function(){
			$(self).hide();
			$(".btn-remove-wishitem").show();
		})
		.fail(function(){ alert("Internal Error!") });
	});

	$(".btn-remove-wishitem").click(function() {
		var self = this;
		var productId = $(this).closest(".show-product").data("product");
		$.post("/wish_item/remove", {pid: productId})
		.done(function(){
			$(self).hide();
			$(".btn-add-wishitem").show();
		})
		.fail(function(){ alert("Internal Error!") });
	});
});