$(function() {
	var errorMsg = "Internal Error!";

	// Add a wish item
	$(".btn-add-wishitem").click(function() {
		var self = this,
			productId = $(this).closest(".show-product").data("product");
		$.post("/wish_item/add", {pid: productId})
		.done(function(){
			$(self).hide();
			$(".btn-remove-wishitem").show();
		})
		.fail(function(){ alert(errorMsg) });
	});

	$(".btn-remove-wishitem").click(function() {
		var self = this,
			productId = $(this).closest(".show-product").data("product");

		$.post("/wish_item/remove", {pid: productId})
		.done(function(){
			$(self).hide();
			$(".btn-add-wishitem").show();
		})
		.fail(function(){ alert(errorMsg) });
	});

	$(".btn-add-ownitem").click(function() {
		var self = this,
			productId = $(this).closest(".show-product").data("product"),
			gotFrom = $(".select-gotFrom").val();

		$.post("/own_item/add", {"ProductId": productId, "GotFrom": gotFrom})
		.done(function(){
			$(".link-add-ownitem-trigger").hide();
			$(".btn-remove-ownitem").show();

			$('#ownitem-modal').modal('hide');
		})
		.fail(function(){ alert(errorMsg) });
	});
});