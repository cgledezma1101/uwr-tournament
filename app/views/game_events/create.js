/*
	Responds to an asynchronous request performed to create an event in a game.

	Expects the following instance variables to be defined:

	@param [GameEvent] game_event The event that has been created
*/
<% if @game_event.persisted? %>
	var event_feeds = $('.event-feed');
	var newEvent = $("<%= j render 'game_events/feed_item', game_event: @game_event %>")
	newEvent.hide().appendTo(event_feeds).fadeIn();
<% else %>
	event_feeds.parent().find('.js-event-text').addClass('chronometer-error');
<% end %>
