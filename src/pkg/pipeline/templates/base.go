package templates

// BaseXML ..
const BaseXML = `
<flow-definition plugin="workflow-job@2.32">
    <actions>
        <io.jenkins.blueocean.service.embedded.BlueOceanUrlAction_-DoNotShowPersistedBlueOceanUrlActions plugin="blueocean-rest-impl@1.16.0"/>
        <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.3.8">
            <jobProperties/>
            <triggers/>
            <parameters/>
            <options/>
        </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
        <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.3.8"/>
    </actions>
    <description>atomci jenkins pipeline</description>
    <keepDependencies>false</keepDependencies>
    <properties>
        <jenkins.model.BuildDiscarderProperty>
            <strategy class="hudson.tasks.LogRotator">
                <daysToKeep>-1</daysToKeep>
                <numToKeep>10</numToKeep>
                <artifactDaysToKeep>-1</artifactDaysToKeep>
                <artifactNumToKeep>-1</artifactNumToKeep>
            </strategy>
        </jenkins.model.BuildDiscarderProperty>
    </properties>
    <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.69">
        <script>
            {{ .Pipeline }}
        </script>
        <sandbox>false</sandbox>
    </definition>
    <triggers/>
    <disabled>false</disabled>
</flow-definition>
`

// DeployBaseXML ..
const DeployBaseXML = `
<flow-definition plugin="workflow-job@2.32">
    <actions>
        <io.jenkins.blueocean.service.embedded.BlueOceanUrlAction_-DoNotShowPersistedBlueOceanUrlActions plugin="blueocean-rest-impl@1.16.0"/>
        <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@1.3.8">
            <jobProperties/>
            <triggers/>
            <parameters/>
            <options/>
        </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
        <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@1.3.8"/>
    </actions>
    <description>deploy pipeline</description>
    <keepDependencies>false</keepDependencies>
    <properties>
        <jenkins.model.BuildDiscarderProperty>
            <strategy class="hudson.tasks.LogRotator">
                <daysToKeep>-1</daysToKeep>
                <numToKeep>60</numToKeep>
                <artifactDaysToKeep>-1</artifactDaysToKeep>
                <artifactNumToKeep>-1</artifactNumToKeep>
            </strategy>
        </jenkins.model.BuildDiscarderProperty>
    </properties>
    <definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@2.69">
        <script>
            {{ .Pipeline }}
        </script>
        <sandbox>false</sandbox>
    </definition>
    <triggers/>
    <disabled>false</disabled>
</flow-definition>
`
