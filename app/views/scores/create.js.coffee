# Dismiss the modal
$('.modal').modal 'hide'

# First, find the player that scored and mark the goal
players_goals = $('#<%= @player.id %>').find('.goals')
players_goals.text(players_goals.text() + 'G')

# Now increment the teams mark
if <%= @player.team_id %> is <%= @game.blue_team_id %>
  selector = '.js-blue-goals'
else
  selector = '.js-white-goals'

teams_score = $(selector).children('h3')
teams_score.text(+(teams_score.text()) + 1)
