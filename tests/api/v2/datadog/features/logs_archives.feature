@v2 @endpoint(logs-archive)
Feature: Logs Archives
  Archives forward all the logs ingested to a cloud storage system.  See the
  [Archives Page](https://app.datadoghq.com/logs/pipelines/archives) for a
  list of the archives currently configured in our UI.

  Background:
    Given a valid "apiKeyAuth" key
    And a valid "appKeyAuth" key
    And an instance of "LogsArchives" API

  @todo
  Scenario: Get all archives leading to OK
    Given new "ListLogsArchives" request
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Create an archive leading to OK
    Given new "CreateLogsArchive" request
    And body {}
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Delete an archive leading to OK
    Given new "DeleteLogsArchive" request
    When the request is sent
    Then the status is 204 OK

  @todo
  Scenario: Get an archive leading to OK
    Given new "GetLogsArchive" request
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Update an archive leading to OK
    Given new "UpdateLogsArchive" request
    And body {}
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Revoke role from an archive leading to OK
    Given new "RemoveRoleFromArchive" request
    And body {}
    When the request is sent
    Then the status is 204 OK

  @todo
  Scenario: List read roles for an archive leading to OK
    Given new "ListArchiveReadRoles" request
    When the request is sent
    Then the status is 200 OK

  @todo
  Scenario: Grant role to an archive leading to OK
    Given new "AddReadRoleToArchive" request
    And body {}
    When the request is sent
    Then the status is 204 OK