/*
  Responds to an asynchronous request to destroy an event from a game

  Expects the following instance variables ot be defined

  @param [GameEvent] game_event The event that has been destroyed
*/
<% if !@game_event.persisted? %>
  $('.game-event#<%= @game_event.id %>').fadeOut();
<% end %>
