Feature: Service listing
  As an authenticated user I can list available services.

  Background:
    Given a running Home Assistant instance
    And an authenticated user

  Scenario: List all services
    When I list all services
    Then I should receive a list of services
    And each service should have a domain and name

  Scenario: Filter services by domain
    When I list services with domain "light"
    Then all services should belong to domain "light"
