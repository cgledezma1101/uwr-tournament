/*
  Responds to the asynchronous action of creating a score in a game

  Expects the following instance variables to be defined

  @param [Player] player Player that scored
  @param [Integer] player_goals Total goals that the player has scored in the match
  @param [Game] game The game where the goal was scored
  @param [Integer] team_goals amount of scores the player's team has in this game
*/

// Dismiss the modal
$('.modal').modal('hide');

// First, find the player that scored and mark the goal
var players_goals = $('#goals-<%= @player.id %>');
players_goals.text(<%= @player_goals %>);
players_goals.parent().removeClass('player-no-scores');

// Now increment the teams mark
<% if @player.team_id == @game.blue_team_id %>
  var selector = '.js-blue-goals';
<% else %>
  var selector = '.js-white-goals';
<% end %>

var teams_score = $(selector);
teams_score.text(<%= @team_goals %>);

addToFeeds("<%= t('score.feed_text', player: @player.name, team: @player.team.name ) %>");
