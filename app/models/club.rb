class Club < ActiveRecord::Base
  has_many :user_clubs, dependent: :destroy
  has_many :users, through: :user_clubs

  has_many :club_admins, dependent: :destroy
  has_many :admins, through: :club_admins, source: :user

  has_many :teams, dependent: :destroy

  # This validation needs to be suppressed until the bug with creating associations using '<<' is fixed
  # validate :has_one_admin
  validates :name, presence: true, uniqueness: true

  # Retrieves all of the members of the club
  def members
    (self.users.to_a + self.admins.to_a).uniq
  end

  private
    def has_one_admin
      errors.add(:admins, I18n.t('club.errors.no_admin')) if self.admins.blank?
    end
end
