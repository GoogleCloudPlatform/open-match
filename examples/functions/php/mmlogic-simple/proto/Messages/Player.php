<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/protobuf-spec/messages.proto

namespace Messages;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Data structure to hold details about a player
 *
 * Generated from protobuf message <code>messages.Player</code>
 */
class Player extends \Google\Protobuf\Internal\Message
{
    /**
     * By convention, a UUID 
     *
     * Generated from protobuf field <code>string id = 1;</code>
     */
    private $id = '';
    /**
     * By convention, a JSON-encoded string
     *
     * Generated from protobuf field <code>string properties = 2;</code>
     */
    private $properties = '';
    /**
     * Optionally used to specify the PlayerPool in which to find a player. 
     *
     * Generated from protobuf field <code>string pool = 3;</code>
     */
    private $pool = '';
    /**
     * Attributes of this player.
     *
     * Generated from protobuf field <code>repeated .messages.Player.Attribute attributes = 4;</code>
     */
    private $attributes;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *           By convention, a UUID 
     *     @type string $properties
     *           By convention, a JSON-encoded string
     *     @type string $pool
     *           Optionally used to specify the PlayerPool in which to find a player. 
     *     @type \Messages\Player\Attribute[]|\Google\Protobuf\Internal\RepeatedField $attributes
     *           Attributes of this player.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Api\ProtobufSpec\Messages::initOnce();
        parent::__construct($data);
    }

    /**
     * By convention, a UUID 
     *
     * Generated from protobuf field <code>string id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * By convention, a UUID 
     *
     * Generated from protobuf field <code>string id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * By convention, a JSON-encoded string
     *
     * Generated from protobuf field <code>string properties = 2;</code>
     * @return string
     */
    public function getProperties()
    {
        return $this->properties;
    }

    /**
     * By convention, a JSON-encoded string
     *
     * Generated from protobuf field <code>string properties = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setProperties($var)
    {
        GPBUtil::checkString($var, True);
        $this->properties = $var;

        return $this;
    }

    /**
     * Optionally used to specify the PlayerPool in which to find a player. 
     *
     * Generated from protobuf field <code>string pool = 3;</code>
     * @return string
     */
    public function getPool()
    {
        return $this->pool;
    }

    /**
     * Optionally used to specify the PlayerPool in which to find a player. 
     *
     * Generated from protobuf field <code>string pool = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPool($var)
    {
        GPBUtil::checkString($var, True);
        $this->pool = $var;

        return $this;
    }

    /**
     * Attributes of this player.
     *
     * Generated from protobuf field <code>repeated .messages.Player.Attribute attributes = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAttributes()
    {
        return $this->attributes;
    }

    /**
     * Attributes of this player.
     *
     * Generated from protobuf field <code>repeated .messages.Player.Attribute attributes = 4;</code>
     * @param \Messages\Player\Attribute[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAttributes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Messages\Player\Attribute::class);
        $this->attributes = $arr;

        return $this;
    }

}

