<template>
    <div class="game-state-bar">
        <span v-if="currentActionTimePositive">
            <span v-format-ns-duration="state.currentActionTimeRemaining"
                  v-help-text="'Remaining time until lack of progress'">
            </span>
            |
            </span>
        <span v-help-text="'Next command'">
            <span v-if="state.nextCommand"
                  :class="teamColorClassNextCommand">
            {{state.nextCommand.type}}
            </span>
            <span v-else>-</span>
        </span>
        |
        <span v-help-text="'Current command'"
              :class="teamColorClassCommand">
            {{state.command.type}}
        </span>
        |
        <span v-help-text="'The current stage of the game'">
            {{stageText}}
        </span>
        |
        <span v-help-text="'Goals for yellow'"
              class="team-yellow">
            {{teamStateYellow.goals}}
        </span>
        :
        <span v-help-text="'Goals for blue'"
              class="team-blue">
            {{teamStateBlue.goals}}
        </span>
        |
        <span v-format-ns-duration="state.stageTimeElapsed"
              v-help-text="'Total time elapsed in the current stage'">
        </span>
        |
        <span v-format-ns-duration="state.stageTimeLeft"
              v-help-text="'Total time left for this stage'">
        </span>
        |
        <span v-format-ns-duration="matchDuration"
              v-help-text="'Total real time elapsed since the match has been started'">
        </span>
    </div>
</template>

<script>
    import {stageNames, TEAM_YELLOW} from "../refereeState";
    import {TEAM_BLUE} from "../refereeState";

    export default {
        name: 'GameStateBar',
        data() {
            return {
                now: Date.now()
            }
        },
        created () {
            let self = this;
            setInterval(function () {
                self.now = Date.now()
            }, 1000)
        },
        computed: {
            state() {
                return this.$store.state.matchState
            },
            matchDuration() {
                let start = Date.parse(this.state.matchTimeStart);
                if (start === 0) {
                    return 0;
                }
                return (this.now - start) * 1000000;
            },
            teamStateYellow() {
                return this.$store.state.matchState.teamState[TEAM_YELLOW]
            },
            teamStateBlue() {
                return this.$store.state.matchState.teamState[TEAM_BLUE]
            },
            teamColorClassNextCommand() {
                return {
                    'team-blue': this.state.nextCommand && this.state.nextCommand.forTeam === TEAM_BLUE,
                    'team-yellow': this.state.nextCommand && this.state.nextCommand.forTeam === TEAM_YELLOW
                }
            },
            teamColorClassCommand() {
                return {
                    'team-blue': this.state.command.forTeam === TEAM_BLUE,
                    'team-yellow': this.state.command.forTeam === TEAM_YELLOW
                }
            },
            stageText() {
                return stageNames[this.state.stage];
            },
            currentActionTimePositive() {
                return !this.state.currentActionTimeRemaining.toString().startsWith('-');
            }
        }
    }
</script>

<style scoped>
    .game-state-bar {
        font-size: 1.2em;
        font-weight: bold;
    }
</style>
