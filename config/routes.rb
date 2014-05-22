UwrTournament::Application.routes.draw do
  # You can have the root of your site routed with "root"
  root 'game#new'

  resources :games
end
