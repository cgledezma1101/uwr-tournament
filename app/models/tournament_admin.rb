class TournamentAdmin < ApplicationRecord
	belongs_to :tournament
	belongs_to :user

	validates :tournament, uniqueness: { scope: :user }
end
