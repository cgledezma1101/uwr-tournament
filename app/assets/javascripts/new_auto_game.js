$(document).on('change', '#newAutoGame .js-team-calculation-param', function(event) {
	event.preventDefault();
	event.stopPropagation();

	var $stageContainer = $(event.currentTarget).parents('.stage-container');
	var stageParam = $stageContainer.find('.js-stage-param').val();
	var positionParam = +($stageContainer.find('.js-position-param').val()) - 1;

	var strategyString = "l|" + positionParam + "|" + stageParam;
	$($stageContainer.data('target-calculation')).attr('value', strategyString);
});

$(document).on('change', '#newAutoGameResult .js-team-calculation-param', function(event) {
	event.preventDefault();
	event.stopPropagation();

	var $stageContainer = $(event.currentTarget).parents('.stage-container');
	var gameParam = $stageContainer.find('.js-game-param').val();
	var outcomeParam = $stageContainer.find('.js-outcome-param').val();

	var strategyString = "g|" + gameParam + "|" + outcomeParam;
	$($stageContainer.data('target-calculation')).attr('value', strategyString);
});