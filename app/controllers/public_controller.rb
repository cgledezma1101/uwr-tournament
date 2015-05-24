class PublicController < ApplicationController
  before_action :authenticate_user!

  # GET /home
  #
  # Redirects to the configured homw page
  def home
    redirect_to user_path(current_user)
  end

  # GET /statistics
  #
  # Shows the statistics of the tournament so far
  def statistics
    @teams = Team.all.to_a.sort{ |a, b| a.points <=> b.points }.reverse
    top_score = Player.includes(:scores)
                      .max{ |a, b| a.scores.count <=> b.scores.count }
                      .scores
                      .count
    @top_scorers = Player.includes(:scores)
                         .select{ |player| player.scores.count == top_score }

    least_goals = Team.all
                      .min{ |a, b| a.goals_against <=> b.goals_against }
                      .goals_against
    @least_goaled = Team.all
                        .select{ |team| team.goals_against == least_goals }
  end
end
