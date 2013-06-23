$(function() {

	// Add a wish item
	$(".btn-wish-item").click(function() {
		var productId = $(this).closest(".show-product").data("product");
		$.post("/wish_item/add", {pid: productId})
		.done(function(){ alert("Done!") })
		.fail(function(){ alert("Fail!") });
	});
});