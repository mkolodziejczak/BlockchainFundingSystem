pragma solidity ^0.8.0;

contract Project {

    enum State {
        Fundraising,
        Expired,
        Failed,
        Milestone1,
        Milestone2,
        Milestone3,
        Ended
    }

    address payable public owner;
    uint public goalBacking;
    uint public goalDeadline;
    uint public milestone1Date;
    uint public milestone2Date;
    uint public milestone3Date;
    uint256 public currentBalance;
    uint256 public totalBacking;
    State public state = State.Fundraising;

    mapping ( address => uint256 ) backing;
    address[] backers;
    address[] firstAcceptanceVote;
    address[] secondAcceptanceVote;
    mapping ( address => bool ) backersToRefund;

    event BackingReceived( address contributor, uint amount, uint totalBacking );
    event OwnerPaid( address recipient, uint percentage, uint256 amount );
    event RefundPaid( address recipient, uint percentage, uint256 amount );
    event AcceptanceVoteCast( address contributor, State state );

    modifier eligibleForFirstAcceptanceVote() {
        require( state == State.Milestone1 && block.timestamp < milestone2Date );
        _;
    }

    modifier eligibleForSecondAcceptanceVote() {
        require( state == State.Milestone2 && block.timestamp < milestone3Date );
        _;
    }

    modifier isFundraising() {
        require( state == State.Fundraising );
        _;
    }

    modifier isFailedOrExpired() {
        require( state == State.Failed || state == State.Expired );
        _;
    }

    constructor
    (
        address payable projectCreator,
        uint milestone1Deadline,
        uint milestone2Deadline,
        uint milestone3Deadline,
        uint goalAmount,
        uint goalDate
    ) {
        owner = projectCreator;
        goalBacking = goalAmount;
        goalDeadline = goalDate;
        milestone1Date = milestone1Deadline;
        milestone2Date = milestone2Deadline;
        milestone3Date = milestone3Deadline;
        currentBalance = 0;
        totalBacking = 0;
    }

    function contribute() external isFundraising payable {
        require( msg.sender != owner );
        backing[ msg.sender ] = backing[ msg.sender ] + msg.value ;
        backers.push( msg.sender );
        backersToRefund[ msg.sender ] = true;
        currentBalance = currentBalance + msg.value ;
        totalBacking = currentBalance;
        emit BackingReceived( msg.sender, msg.value, currentBalance );
    }

    function payOut( uint percentage ) internal returns ( bool ) {
        uint256 toPayOut = totalBacking * ( percentage / 100 );

        if ( owner.send( toPayOut ) ) {
            emit OwnerPaid( owner, percentage, toPayOut );
            currentBalance = currentBalance - toPayOut;
            return true;
        }
        return false;
    }

    function getRefund( uint percentage ) internal returns ( bool ) {
        bool failure = false;

        for ( uint i = 0; i < backers.length; i++ ) {
            uint256 amountToRefund = backing[ backers[i] ] * ( percentage / 100 );

            if( backersToRefund[ backers[i] ] ) {
                if( payable( backers[i] ).send( amountToRefund ) ) {
                    emit RefundPaid( backers[i], percentage, amountToRefund );
                    backersToRefund[ backers[i] ] = false;
                } else {
                    failure = true;
                }
            }
        }

        return failure;
    }


    function voteForMilestone1Acceptance() external eligibleForFirstAcceptanceVote {
        require( backing[ msg.sender ] > 0 );
        castVote( firstAcceptanceVote, msg.sender, State.Milestone2 );
        emit AcceptanceVoteCast( msg.sender, State.Milestone1 );
    }

    function voteForMilestone2Acceptance() external eligibleForSecondAcceptanceVote {
        require( backing[ msg.sender ] > 0 );
        castVote( secondAcceptanceVote, msg.sender, State.Milestone3 );
        emit AcceptanceVoteCast( msg.sender, State.Milestone2 );
    }

    function castVote( address[] storage votes, address sender, State nextState ) internal {
        votes.push( sender );
        checkIfVotePassed( votes, nextState );
    }

    function checkIfVotePassed( address[] storage votes , State nextState ) internal {
        uint256 sum = 0;
        for ( uint i = 0; i < votes.length; i++ ) {
            sum = sum + backing[ votes[i] ];
        }
        uint256 percentage = ( sum / totalBacking ) * 100;

        if( percentage > 70 ) {
            proceedToNextMilestone( nextState );
        }
    }

    function getDetails() public view returns
    (
        address payable projectStarter,
        uint256 deadline,
        State currentState,
        uint256 currentAmount,
        uint256 amountGoal,
        uint256 totalAmount
    ) {
        projectStarter = owner;
        deadline = goalDeadline;
        currentState = state;
        currentAmount = currentBalance;
        totalAmount = totalBacking;
        amountGoal = goalBacking;
    }

    function processProjectStatus() external {
        if( state == State.Fundraising && block.timestamp > goalDeadline ) {
            if ( totalBacking >= goalBacking ) {
                proceedToNextMilestone( State.Milestone1 );
            } else {
                state = State.Expired;
                getRefund( 100 );
            }
        } else if ( state == State.Milestone1 && block.timestamp > milestone1Date ) {
            state = State.Failed;
            getRefund( 70 );
        } else if ( state == State.Milestone2 && block.timestamp > milestone2Date ) {
            state = State.Failed;
            getRefund( 30 );
        } else if ( state == State.Milestone3 && block.timestamp > milestone3Date ) {
            state = State.Ended;
        }
    }

    function proceedToNextMilestone( State nextState ) internal {
        if( nextState == State.Milestone1 ) {
            if ( payOut( 30 ) ) {
                state = State.Milestone1;
            }
        } else if( nextState == State.Milestone2 ) {
            if ( payOut( 40 ) ) {
                state = State.Milestone2;
            }
        } else if ( nextState == State.Milestone3 ) {
            if ( payOut( 30 ) ) {
                state = State.Milestone3;
            }
        }
    }

    function checkIfFullyRefunded() external view isFailedOrExpired returns (bool) {
        bool refunded = true;
        for ( uint i = 0; i < backers.length; i++ ) {
            if( backersToRefund[ backers[i] ] ) {
                refunded = false;
            }
        }
        return refunded;
    }
}
