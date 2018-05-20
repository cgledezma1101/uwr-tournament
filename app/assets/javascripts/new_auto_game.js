$(document).on('change', '#newAutoGame .js-team-calculation-param', function(event) {
	event.preventDefault();
	event.stopPropagation();

	var $stageContainer = $(event.currentTarget).parents('.stage-container');
	var stageParam = $stageContainer.find('.js-stage-param').val();
	var positionParam = +($stageContainer.find('.js-position-param').val()) - 1;

	var strategyString = "l|" + positionParam + "|" + stageParam;
	$($stageContainer.data('target-calculation')).attr('value', strategyString);
})