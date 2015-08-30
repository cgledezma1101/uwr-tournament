class TournamentInvitation < ActiveRecord::Base
  belongs_to :tournament
  belongs_to :club
end
