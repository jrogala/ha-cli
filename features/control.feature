Feature: Entity control
  As an authenticated user I can turn entities on, off, or toggle them.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: Turn on an entity
    Given entity "light.kitchen" exists
    When I turn on "light.kitchen"
    Then the action should succeed with action "on"

  Scenario: Turn off an entity
    Given entity "light.kitchen" exists
    When I turn off "light.kitchen"
    Then the action should succeed with action "off"

  Scenario: Toggle an entity
    Given entity "light.kitchen" exists
    When I toggle "light.kitchen"
    Then the action should succeed with action "toggle"
