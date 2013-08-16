$(function() {
	var errorMsg = "Internal Error!";

	// Add product to wish item
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

	// Remove a wish item
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

	// Add product to own item
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

	// Remove a own item
	$(".btn-remove-ownitem").click(function() {
		var self = this,
			productId = $(this).closest(".show-product").data("product");

		$.post("/own_item/remove", {pid: productId})
		.done(function(){
			$(self).hide();
			$(".link-add-ownitem-trigger").show();
		})
		.fail(function(){ alert(errorMsg) });
	});

	// Add/Edit product > select Category
	$(".category-select").change(function(){
		var categoryId = $(this).val();
		$(".subcategory-efficacy").hide();
		$(".category-"+categoryId).show();

	});

});