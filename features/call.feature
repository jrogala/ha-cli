Feature: Service call
  As an authenticated user I can call a Home Assistant service.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: Call a service
    When I call service "light" "turn_on" with data:
      | entity_id     |
      | light.kitchen |
    Then the service call should succeed
