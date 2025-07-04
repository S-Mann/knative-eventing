<p>Packages:</p>
<ul>
<li>
<a href="#duck.knative.dev%2fv1">duck.knative.dev/v1</a>
</li>
<li>
<a href="#duck.knative.dev%2fv1alpha1">duck.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#duck.knative.dev%2fv1beta1">duck.knative.dev/v1beta1</a>
</li>
<li>
<a href="#eventing.knative.dev%2fv1">eventing.knative.dev/v1</a>
</li>
<li>
<a href="#eventing.knative.dev%2fv1alpha1">eventing.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#eventing.knative.dev%2fv1beta1">eventing.knative.dev/v1beta1</a>
</li>
<li>
<a href="#eventing.knative.dev%2fv1beta2">eventing.knative.dev/v1beta2</a>
</li>
<li>
<a href="#eventing.knative.dev%2fv1beta3">eventing.knative.dev/v1beta3</a>
</li>
<li>
<a href="#flows.knative.dev%2fv1">flows.knative.dev/v1</a>
</li>
<li>
<a href="#messaging.knative.dev%2fv1">messaging.knative.dev/v1</a>
</li>
<li>
<a href="#sinks.knative.dev%2fv1alpha1">sinks.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#sources.knative.dev%2fv1">sources.knative.dev/v1</a>
</li>
<li>
<a href="#sources.knative.dev%2fv1alpha1">sources.knative.dev/v1alpha1</a>
</li>
<li>
<a href="#sources.knative.dev%2fv1beta2">sources.knative.dev/v1beta2</a>
</li>
</ul>
<h2 id="duck.knative.dev/v1">duck.knative.dev/v1</h2>
<p>
<p>Package v1 is the v1 version of the API.</p>
</p>
Resource Types:
<ul></ul>
<h3 id="duck.knative.dev/v1.AppliedEventPoliciesStatus">AppliedEventPoliciesStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.ChannelableStatus">ChannelableStatus</a>, <a href="#eventing.knative.dev/v1.BrokerStatus">BrokerStatus</a>, <a href="#eventing.knative.dev/v1alpha1.RequestReplyStatus">RequestReplyStatus</a>, <a href="#flows.knative.dev/v1.ParallelStatus">ParallelStatus</a>, <a href="#flows.knative.dev/v1.SequenceStatus">SequenceStatus</a>, <a href="#sinks.knative.dev/v1alpha1.IntegrationSinkStatus">IntegrationSinkStatus</a>, <a href="#sinks.knative.dev/v1alpha1.JobSinkStatus">JobSinkStatus</a>)
</p>
<p>
<p>AppliedEventPoliciesStatus contains the list of policies which apply to a resource.
This type is intended to be embedded into a status struct.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>policies</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPolicyRef">
[]AppliedEventPolicyRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Policies holds the list of applied EventPolicies</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.AppliedEventPolicyRef">AppliedEventPolicyRef
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">AppliedEventPoliciesStatus</a>)
</p>
<p>
<p>AppliedEventPolicyRef is the reference to an EventPolicy</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion of the applied EventPolicy.
This indicates, which version of EventPolicy is supported by the resource.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the applied EventPolicy</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.BackoffPolicyType">BackoffPolicyType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.DeliverySpec">DeliverySpec</a>)
</p>
<p>
<p>BackoffPolicyType is the type for backoff policies</p>
</p>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;exponential&#34;</p></td>
<td><p>Exponential backoff policy</p>
</td>
</tr><tr><td><p>&#34;linear&#34;</p></td>
<td><p>Linear backoff policy</p>
</td>
</tr></tbody>
</table>
<h3 id="duck.knative.dev/v1.Channelable">Channelable
</h3>
<p>
<p>Channelable is a skeleton type wrapping Subscribable and Addressable in the manner we expect resource writers
defining compatible resources to embed it. We will typically use this type to deserialize
Channelable ObjectReferences and access their subscription and address data.  This is not a real resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>Spec is the part where the Channelable fulfills the Subscribable contract.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>SubscribableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains the default delivery spec for each subscription
to this Channelable. Each subscription delivery spec, if any, overrides this
global delivery spec.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableStatus">
ChannelableStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.ChannelableSpec">ChannelableSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Channelable">Channelable</a>, <a href="#messaging.knative.dev/v1.ChannelSpec">ChannelSpec</a>, <a href="#messaging.knative.dev/v1.InMemoryChannelSpec">InMemoryChannelSpec</a>)
</p>
<p>
<p>ChannelableSpec contains Spec of the Channelable object</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SubscribableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains the default delivery spec for each subscription
to this Channelable. Each subscription delivery spec, if any, overrides this
global delivery spec.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.ChannelableStatus">ChannelableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.Channelable">Channelable</a>, <a href="#messaging.knative.dev/v1.ChannelStatus">ChannelStatus</a>, <a href="#messaging.knative.dev/v1.InMemoryChannelStatus">InMemoryChannelStatus</a>)
</p>
<p>
<p>ChannelableStatus contains the Status of a Channelable object.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the Channelable fulfills the Addressable contract.</p>
</td>
</tr>
<tr>
<td>
<code>SubscribableStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscribableStatus">
SubscribableStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableStatus</code> are embedded into this type.)
</p>
<p>Subscribers is populated with the statuses of each of the Channelable&rsquo;s subscribers.</p>
</td>
</tr>
<tr>
<td>
<code>DeliveryStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliveryStatus">
DeliveryStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>DeliveryStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>DeliveryStatus contains a resolved URL to the dead letter sink address, and any other
resolved delivery options.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this Channel</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.DeliverySpec">DeliverySpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.ChannelableSpec">ChannelableSpec</a>, <a href="#duck.knative.dev/v1.SubscriberSpec">SubscriberSpec</a>, <a href="#eventing.knative.dev/v1.BrokerSpec">BrokerSpec</a>, <a href="#eventing.knative.dev/v1.TriggerSpec">TriggerSpec</a>, <a href="#eventing.knative.dev/v1alpha1.RequestReplySpec">RequestReplySpec</a>, <a href="#flows.knative.dev/v1.ParallelBranch">ParallelBranch</a>, <a href="#flows.knative.dev/v1.SequenceStep">SequenceStep</a>, <a href="#messaging.knative.dev/v1.SubscriptionSpec">SubscriptionSpec</a>)
</p>
<p>
<p>DeliverySpec contains the delivery options for event senders,
such as channelable and source.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deadLetterSink</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterSink is the sink receiving event that could not be sent to
a destination.</p>
</td>
</tr>
<tr>
<td>
<code>retry</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Retry is the minimum number of retries the sender should attempt when
sending an event before moving it to the dead letter sink.</p>
</td>
</tr>
<tr>
<td>
<code>timeout</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Timeout is the timeout of each single request. The value must be greater than 0.
More information on Duration format:
- <a href="https://www.iso.org/iso-8601-date-and-time-format.html">https://www.iso.org/iso-8601-date-and-time-format.html</a>
- <a href="https://en.wikipedia.org/wiki/ISO_8601">https://en.wikipedia.org/wiki/ISO_8601</a></p>
<p>Note: This API is EXPERIMENTAL and might break anytime. For more details: <a href="https://github.com/knative/eventing/issues/5148">https://github.com/knative/eventing/issues/5148</a></p>
</td>
</tr>
<tr>
<td>
<code>backoffPolicy</code><br/>
<em>
<a href="#duck.knative.dev/v1.BackoffPolicyType">
BackoffPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>BackoffPolicy is the retry backoff policy (linear, exponential).</p>
</td>
</tr>
<tr>
<td>
<code>backoffDelay</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>BackoffDelay is the delay before retrying.
More information on Duration format:
- <a href="https://www.iso.org/iso-8601-date-and-time-format.html">https://www.iso.org/iso-8601-date-and-time-format.html</a>
- <a href="https://en.wikipedia.org/wiki/ISO_8601">https://en.wikipedia.org/wiki/ISO_8601</a></p>
<p>For linear policy, backoff delay is backoffDelay*<numberOfRetries>.
For exponential policy, backoff delay is backoffDelay*2^<numberOfRetries>.</p>
</td>
</tr>
<tr>
<td>
<code>retryAfterMax</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>RetryAfterMax provides an optional upper bound on the duration specified in a &ldquo;Retry-After&rdquo; header
when calculating backoff times for retrying 429 and 503 response codes.  Setting the value to
zero (&ldquo;PT0S&rdquo;) can be used to opt-out of respecting &ldquo;Retry-After&rdquo; header values altogether. This
value only takes effect if &ldquo;Retry&rdquo; is configured, and also depends on specific implementations
(Channels, Sources, etc.) choosing to provide this capability.</p>
<p>Note: This API is EXPERIMENTAL and might be changed at anytime. While this experimental
feature is in the Alpha/Beta stage, you must provide a valid value to opt-in for
supporting &ldquo;Retry-After&rdquo; headers.  When the feature becomes Stable/GA &ldquo;Retry-After&rdquo;
headers will be respected by default, and you can choose to specify &ldquo;PT0S&rdquo; to
opt-out of supporting &ldquo;Retry-After&rdquo; headers.
For more details: <a href="https://github.com/knative/eventing/issues/5811">https://github.com/knative/eventing/issues/5811</a></p>
<p>More information on Duration format:
- <a href="https://www.iso.org/iso-8601-date-and-time-format.html">https://www.iso.org/iso-8601-date-and-time-format.html</a>
- <a href="https://en.wikipedia.org/wiki/ISO_8601">https://en.wikipedia.org/wiki/ISO_8601</a></p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#duck.knative.dev/v1.FormatType">
FormatType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>format specifies the desired event format for the cloud event.
It can be one of the following values:
- nil: default value, no specific format required.
- &ldquo;JSON&rdquo;: indicates the event should be in structured mode.
- &ldquo;binary&rdquo;: indicates the event should be in binary mode.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.DeliveryStatus">DeliveryStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.ChannelableStatus">ChannelableStatus</a>, <a href="#eventing.knative.dev/v1.BrokerStatus">BrokerStatus</a>, <a href="#eventing.knative.dev/v1.TriggerStatus">TriggerStatus</a>, <a href="#messaging.knative.dev/v1.SubscriptionStatusPhysicalSubscription">SubscriptionStatusPhysicalSubscription</a>)
</p>
<p>
<p>DeliveryStatus contains the Status of an object supporting delivery options. This type is intended to be embedded into a status struct.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deadLetterSinkUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterSink is a KReference that is the reference to the native, platform specific channel
where failed events are sent to.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterSinkCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterSinkCACerts are Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a>.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterSinkAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterSinkAudience is the OIDC audience of the DeadLetterSink</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.FormatType">FormatType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.DeliverySpec">DeliverySpec</a>)
</p>
<p>
<p>FormatType is the type for delivery format</p>
</p>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;binary&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;json&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="duck.knative.dev/v1.Subscribable">Subscribable
</h3>
<p>
<p>Subscribable is a skeleton type wrapping Subscribable in the manner we expect resource writers
defining compatible resources to embed it. We will typically use this type to deserialize
SubscribableType ObjectReferences and access the Subscription data.  This is not a real resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>SubscribableSpec is the part where Subscribable object is
configured as to be compatible with Subscribable contract.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscriberSpec">
[]SubscriberSpec
</a>
</em>
</td>
<td>
<p>This is the list of subscriptions for this subscribable.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscribableStatus">
SubscribableStatus
</a>
</em>
</td>
<td>
<p>SubscribableStatus is the part where SubscribableStatus object is
configured as to be compatible with Subscribable contract.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SubscribableSpec">SubscribableSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.ChannelableSpec">ChannelableSpec</a>, <a href="#duck.knative.dev/v1.Subscribable">Subscribable</a>)
</p>
<p>
<p>SubscribableSpec shows how we expect folks to embed Subscribable in their Spec field.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscriberSpec">
[]SubscriberSpec
</a>
</em>
</td>
<td>
<p>This is the list of subscriptions for this subscribable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SubscribableStatus">SubscribableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.ChannelableStatus">ChannelableStatus</a>, <a href="#duck.knative.dev/v1.Subscribable">Subscribable</a>)
</p>
<p>
<p>SubscribableStatus is the schema for the subscribable&rsquo;s status portion of the status
section of the resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1.SubscriberStatus">
[]SubscriberStatus
</a>
</em>
</td>
<td>
<p>This is the list of subscription&rsquo;s statuses for this channel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SubscriberSpec">SubscriberSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.SubscribableSpec">SubscribableSpec</a>)
</p>
<p>
<p>SubscriberSpec defines a single subscriber to a Subscribable.</p>
<p>At least one of SubscriberURI and ReplyURI must be present</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name is used to identify the original subscription object.</p>
</td>
</tr>
<tr>
<td>
<code>uid</code><br/>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UID is used to understand the origin of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>generation</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Generation of the origin of the subscriber with uid:UID.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberURI is the endpoint for the subscriber</p>
</td>
</tr>
<tr>
<td>
<code>subscriberCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberCACerts is the Certification Authority (CA) certificates in PEM
format according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a> for the
subscriberUri</p>
</td>
</tr>
<tr>
<td>
<code>subscriberAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberAudience is the OIDC audience for the subscriberUri.</p>
</td>
</tr>
<tr>
<td>
<code>replyUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyURI is the endpoint for the reply</p>
</td>
</tr>
<tr>
<td>
<code>replyCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyCACerts is the Certification Authority (CA) certificates in PEM
format according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a> for the
replyUri.</p>
</td>
</tr>
<tr>
<td>
<code>replyAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyAudience is the OIDC audience for the replyUri.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains options controlling the event delivery</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth contains the service account name for the subscription</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1.SubscriberStatus">SubscriberStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1.SubscribableStatus">SubscribableStatus</a>)
</p>
<p>
<p>SubscriberStatus defines the status of a single subscriber to a Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uid</code><br/>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UID is used to understand the origin of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Generation of the origin of the subscriber with uid:UID.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#conditionstatus-v1-core">
Kubernetes core/v1.ConditionStatus
</a>
</em>
</td>
<td>
<p>Status of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A human readable message indicating details of Ready status.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth provides the relevant information for OIDC authentication.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="duck.knative.dev/v1alpha1">duck.knative.dev/v1alpha1</h2>
<p>
</p>
Resource Types:
<ul></ul>
<h3 id="duck.knative.dev/v1alpha1.Placeable">Placeable
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1alpha1.PlaceableStatus">PlaceableStatus</a>)
</p>
<p>
<p>Placeable is a list of podName and virtual replicas pairs.
Each pair represents the assignment of virtual replicas to a pod</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxAllowedVReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>placements</code><br/>
<em>
<a href="#duck.knative.dev/v1alpha1.Placement">
[]Placement
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1alpha1.PlaceableStatus">PlaceableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1alpha1.PlaceableType">PlaceableType</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Placeable</code><br/>
<em>
<a href="#duck.knative.dev/v1alpha1.Placeable">
Placeable
</a>
</em>
</td>
<td>
<p>
(Members of <code>Placeable</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1alpha1.PlaceableType">PlaceableType
</h3>
<p>
<p>PlaceableType is a skeleton type wrapping Placeable in the manner we expect
resource writers defining compatible resources to embed it.  We will
typically use this type to deserialize Placeable ObjectReferences and
access the Placeable data.  This is not a real resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1alpha1.PlaceableStatus">
PlaceableStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1alpha1.Placement">Placement
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1alpha1.Placeable">Placeable</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>podName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodName is the name of the pod where the resource is placed</p>
</td>
</tr>
<tr>
<td>
<code>vreplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>VReplicas is the number of virtual replicas assigned to in the pod</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="duck.knative.dev/v1beta1">duck.knative.dev/v1beta1</h2>
<p>
<p>Package v1beta1 is the v1beta1 version of the API.</p>
</p>
Resource Types:
<ul></ul>
<h3 id="duck.knative.dev/v1beta1.BackoffPolicyType">BackoffPolicyType
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.DeliverySpec">DeliverySpec</a>)
</p>
<p>
<p>BackoffPolicyType is the type for backoff policies</p>
</p>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;exponential&#34;</p></td>
<td><p>Exponential backoff policy</p>
</td>
</tr><tr><td><p>&#34;linear&#34;</p></td>
<td><p>Linear backoff policy</p>
</td>
</tr></tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Channelable">Channelable
</h3>
<p>
<p>Channelable is a skeleton type wrapping Subscribable and Addressable in the manner we expect resource writers
defining compatible resources to embed it. We will typically use this type to deserialize
Channelable ObjectReferences and access their subscription and address data.  This is not a real resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>Spec is the part where the Channelable fulfills the Subscribable contract.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>SubscribableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains options controlling the event delivery</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.ChannelableStatus">
ChannelableStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.ChannelableSpec">ChannelableSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Channelable">Channelable</a>)
</p>
<p>
<p>ChannelableSpec contains Spec of the Channelable object</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SubscribableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains options controlling the event delivery</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.ChannelableStatus">ChannelableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.Channelable">Channelable</a>)
</p>
<p>
<p>ChannelableStatus contains the Status of a Channelable object.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<p>AddressStatus is the part where the Channelable fulfills the Addressable contract.</p>
</td>
</tr>
<tr>
<td>
<code>SubscribableStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscribableStatus">
SubscribableStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SubscribableStatus</code> are embedded into this type.)
</p>
<p>Subscribers is populated with the statuses of each of the Channelable&rsquo;s subscribers.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterChannel</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterChannel is a KReference and is set by the channel when it supports native error handling via a channel
Failed messages are delivered here.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.DeliverySpec">DeliverySpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.ChannelableSpec">ChannelableSpec</a>, <a href="#duck.knative.dev/v1beta1.SubscriberSpec">SubscriberSpec</a>)
</p>
<p>
<p>DeliverySpec contains the delivery options for event senders,
such as channelable and source.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deadLetterSink</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterSink is the sink receiving event that could not be sent to
a destination.</p>
</td>
</tr>
<tr>
<td>
<code>retry</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Retry is the minimum number of retries the sender should attempt when
sending an event before moving it to the dead letter sink.</p>
</td>
</tr>
<tr>
<td>
<code>timeout</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timeout is the timeout of each single request.
More information on Duration format:
- <a href="https://www.iso.org/iso-8601-date-and-time-format.html">https://www.iso.org/iso-8601-date-and-time-format.html</a>
- <a href="https://en.wikipedia.org/wiki/ISO_8601">https://en.wikipedia.org/wiki/ISO_8601</a></p>
</td>
</tr>
<tr>
<td>
<code>backoffPolicy</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.BackoffPolicyType">
BackoffPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>BackoffPolicy is the retry backoff policy (linear, exponential).</p>
</td>
</tr>
<tr>
<td>
<code>backoffDelay</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>BackoffDelay is the delay before retrying.
More information on Duration format:
- <a href="https://www.iso.org/iso-8601-date-and-time-format.html">https://www.iso.org/iso-8601-date-and-time-format.html</a>
- <a href="https://en.wikipedia.org/wiki/ISO_8601">https://en.wikipedia.org/wiki/ISO_8601</a></p>
<p>For linear policy, backoff delay is backoffDelay*<numberOfRetries>.
For exponential policy, backoff delay is backoffDelay*2^<numberOfRetries>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.DeliveryStatus">DeliveryStatus
</h3>
<p>
<p>DeliveryStatus contains the Status of an object supporting delivery options.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deadLetterChannel</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeadLetterChannel is a KReference that is the reference to the native, platform specific channel
where failed events are sent to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.Subscribable">Subscribable
</h3>
<p>
<p>Subscribable is a skeleton type wrapping Subscribable in the manner we expect resource writers
defining compatible resources to embed it. We will typically use this type to deserialize
SubscribableType ObjectReferences and access the Subscription data.  This is not a real resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscribableSpec">
SubscribableSpec
</a>
</em>
</td>
<td>
<p>SubscribableSpec is the part where Subscribable object is
configured as to be compatible with Subscribable contract.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscriberSpec">
[]SubscriberSpec
</a>
</em>
</td>
<td>
<p>This is the list of subscriptions for this subscribable.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscribableStatus">
SubscribableStatus
</a>
</em>
</td>
<td>
<p>SubscribableStatus is the part where SubscribableStatus object is
configured as to be compatible with Subscribable contract.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SubscribableSpec">SubscribableSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.ChannelableSpec">ChannelableSpec</a>, <a href="#duck.knative.dev/v1beta1.Subscribable">Subscribable</a>)
</p>
<p>
<p>SubscribableSpec shows how we expect folks to embed Subscribable in their Spec field.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscriberSpec">
[]SubscriberSpec
</a>
</em>
</td>
<td>
<p>This is the list of subscriptions for this subscribable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SubscribableStatus">SubscribableStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.ChannelableStatus">ChannelableStatus</a>, <a href="#duck.knative.dev/v1beta1.Subscribable">Subscribable</a>)
</p>
<p>
<p>SubscribableStatus is the schema for the subscribable&rsquo;s status portion of the status
section of the resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscribers</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.SubscriberStatus">
[]SubscriberStatus
</a>
</em>
</td>
<td>
<p>This is the list of subscription&rsquo;s statuses for this channel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SubscriberSpec">SubscriberSpec
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.SubscribableSpec">SubscribableSpec</a>)
</p>
<p>
<p>SubscriberSpec defines a single subscriber to a Subscribable.</p>
<p>At least one of SubscriberURI and ReplyURI must be present</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uid</code><br/>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UID is used to understand the origin of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>generation</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Generation of the origin of the subscriber with uid:UID.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberURI is the endpoint for the subscriber</p>
</td>
</tr>
<tr>
<td>
<code>replyUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyURI is the endpoint for the reply</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1beta1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DeliverySpec contains options controlling the event delivery</p>
</td>
</tr>
</tbody>
</table>
<h3 id="duck.knative.dev/v1beta1.SubscriberStatus">SubscriberStatus
</h3>
<p>
(<em>Appears on:</em><a href="#duck.knative.dev/v1beta1.SubscribableStatus">SubscribableStatus</a>)
</p>
<p>
<p>SubscriberStatus defines the status of a single subscriber to a Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>uid</code><br/>
<em>
<a href="https://godoc.org/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>UID is used to understand the origin of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Generation of the origin of the subscriber with uid:UID.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#conditionstatus-v1-core">
Kubernetes core/v1.ConditionStatus
</a>
</em>
</td>
<td>
<p>Status of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A human readable message indicating details of Ready status.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="eventing.knative.dev/v1">eventing.knative.dev/v1</h2>
<p>
<p>Package v1 is the v1 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#eventing.knative.dev/v1.Broker">Broker</a>
</li><li>
<a href="#eventing.knative.dev/v1.Trigger">Trigger</a>
</li></ul>
<h3 id="eventing.knative.dev/v1.Broker">Broker
</h3>
<p>
<p>Broker collects a pool of events that are consumable using Triggers. Brokers
provide a well-known endpoint for event delivery that senders can use with
minimal knowledge of the event routing strategy. Subscribers use Triggers to
request delivery of events from a Broker&rsquo;s pool to a specific URL or
Addressable endpoint.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Broker</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1.BrokerSpec">
BrokerSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Broker.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>config</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Config is a KReference to the configuration that specifies
configuration options for this Broker. For example, this could be
a pointer to a ConfigMap.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery contains the delivery spec for each trigger
to this Broker. Each trigger delivery spec, if any, overrides this
global delivery spec.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1.BrokerStatus">
BrokerStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Broker. This data may be out of
date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.Trigger">Trigger
</h3>
<p>
<p>Trigger represents a request to have events delivered to a subscriber from a
Broker&rsquo;s event pool.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Trigger</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1.TriggerSpec">
TriggerSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Trigger.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<p>Broker is the broker that this trigger receives events from.</p>
</td>
</tr>
<tr>
<td>
<code>brokerRef</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>BrokerRef is the broker that is used for cross-namespace referencing.</p>
</td>
</tr>
<tr>
<td>
<code>filter</code><br/>
<em>
<a href="#eventing.knative.dev/v1.TriggerFilter">
TriggerFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filter is the filter to apply against all events from the Broker. Only events that pass this
filter will be sent to the Subscriber. If not specified, will default to allowing all events.</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is an experimental field that conforms to the CNCF CloudEvents Subscriptions
API. It&rsquo;s an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event MUST
NOT be sent to the Subscriber. If all the filter expressions in the array
evaluate to true, the event MUST be attempted to be delivered. Absence of
a filter or empty array implies a value of true. In the event of users
specifying both Filter and Filters, then the latter will override the former.
This will allow users to try out the effect of the new Filters field
without compromising the existing attribute-based Filter and try it out on existing
Trigger objects.</p>
</td>
</tr>
<tr>
<td>
<code>subscriber</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>Subscriber is the addressable that receives events from the Broker that pass
the Filter. It is required.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery contains the delivery spec for this specific trigger.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1.TriggerStatus">
TriggerStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Trigger. This data may be out of
date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.BrokerSpec">BrokerSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.Broker">Broker</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>config</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Config is a KReference to the configuration that specifies
configuration options for this Broker. For example, this could be
a pointer to a ConfigMap.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery contains the delivery spec for each trigger
to this Broker. Each trigger delivery spec, if any, overrides this
global delivery spec.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.BrokerStatus">BrokerStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.Broker">Broker</a>)
</p>
<p>
<p>BrokerStatus represents the current state of a Broker.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Broker that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the Broker fulfills the Addressable contract.
It exposes the endpoint as an URI to get events delivered into the Broker mesh.</p>
</td>
</tr>
<tr>
<td>
<code>DeliveryStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliveryStatus">
DeliveryStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>DeliveryStatus</code> are embedded into this type.)
</p>
<p>DeliveryStatus contains a resolved URL to the dead letter sink address, and any other
resolved delivery options.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this Broker</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.SubscriptionsAPIFilter">SubscriptionsAPIFilter
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">SubscriptionsAPIFilter</a>, <a href="#eventing.knative.dev/v1.TriggerSpec">TriggerSpec</a>, <a href="#eventing.knative.dev/v1alpha1.EventPolicySpec">EventPolicySpec</a>, <a href="#sources.knative.dev/v1.ApiServerSourceSpec">ApiServerSourceSpec</a>)
</p>
<p>
<p>SubscriptionsAPIFilter allows defining a filter expression using CloudEvents
Subscriptions API. If multiple filters are specified, then the same semantics
of SubscriptionsAPIFilter.All is applied. If no filter dialect or empty
object is specified, then the filter always accept the events.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>all</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>All evaluates to true if all the nested expressions evaluate to true.
It must contain at least one filter expression.</p>
</td>
</tr>
<tr>
<td>
<code>any</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Any evaluates to true if at least one of the nested expressions evaluates
to true. It must contain at least one filter expression.</p>
</td>
</tr>
<tr>
<td>
<code>not</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Not evaluates to true if the nested expression evaluates to false.</p>
</td>
</tr>
<tr>
<td>
<code>exact</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Exact evaluates to true if the values of the matching CloudEvents attributes MUST
all exactly match with the associated value String specified (case-sensitive).
The keys are the names of the CloudEvents attributes to be matched,
and their values are the String values to use in the comparison.
The attribute name and value specified in the filter express MUST NOT be
empty strings.</p>
</td>
</tr>
<tr>
<td>
<code>prefix</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Prefix evaluates to true if the values of the matching CloudEvents attributes MUST
all start with the associated value String specified (case sensitive).
The keys are the names of the CloudEvents attributes to be matched,
and their values are the String values to use in the comparison.
The attribute name and value specified in the filter express MUST NOT be
empty strings.</p>
</td>
</tr>
<tr>
<td>
<code>suffix</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Suffix evaluates to true if the values of the matching CloudEvents attributes MUST
all end with the associated value String specified (case sensitive).
The keys are the names of the CloudEvents attributes to be matched,
and their values are the String values to use in the comparison.
The attribute name and value specified in the filter express MUST NOT be
empty strings.</p>
</td>
</tr>
<tr>
<td>
<code>cesql</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>CESQL is a CloudEvents SQL expression that will be evaluated to true or false against each CloudEvent.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.TriggerFilter">TriggerFilter
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.TriggerSpec">TriggerSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>attributes</code><br/>
<em>
<a href="#eventing.knative.dev/v1.TriggerFilterAttributes">
TriggerFilterAttributes
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Attributes filters events by exact match on event context attributes.
Each key in the map is compared with the equivalent key in the event
context. An event passes the filter if all values are equal to the
specified values. Nested context attributes are not supported as keys. Only
string values are supported.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.TriggerFilterAttributes">TriggerFilterAttributes
(<code>map[string]string</code> alias)</p></h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.TriggerFilter">TriggerFilter</a>)
</p>
<p>
<p>TriggerFilterAttributes is a map of context attribute names to values for
filtering by equality. Only exact matches will pass the filter. You can use
the value ” to indicate all strings match.</p>
</p>
<h3 id="eventing.knative.dev/v1.TriggerSpec">TriggerSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.Trigger">Trigger</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<p>Broker is the broker that this trigger receives events from.</p>
</td>
</tr>
<tr>
<td>
<code>brokerRef</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>BrokerRef is the broker that is used for cross-namespace referencing.</p>
</td>
</tr>
<tr>
<td>
<code>filter</code><br/>
<em>
<a href="#eventing.knative.dev/v1.TriggerFilter">
TriggerFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filter is the filter to apply against all events from the Broker. Only events that pass this
filter will be sent to the Subscriber. If not specified, will default to allowing all events.</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is an experimental field that conforms to the CNCF CloudEvents Subscriptions
API. It&rsquo;s an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event MUST
NOT be sent to the Subscriber. If all the filter expressions in the array
evaluate to true, the event MUST be attempted to be delivered. Absence of
a filter or empty array implies a value of true. In the event of users
specifying both Filter and Filters, then the latter will override the former.
This will allow users to try out the effect of the new Filters field
without compromising the existing attribute-based Filter and try it out on existing
Trigger objects.</p>
</td>
</tr>
<tr>
<td>
<code>subscriber</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>Subscriber is the addressable that receives events from the Broker that pass
the Filter. It is required.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery contains the delivery spec for this specific trigger.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1.TriggerStatus">TriggerStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1.Trigger">Trigger</a>)
</p>
<p>
<p>TriggerStatus represents the current state of a Trigger.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Trigger that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberURI is the resolved URI of the receiver for this Trigger.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberCACerts is the Certification Authority (CA) certificates in PEM format
according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a> of the receiver for this Trigger.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberAudience is the OIDC audience of the subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>DeliveryStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliveryStatus">
DeliveryStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>DeliveryStatus</code> are embedded into this type.)
</p>
<p>DeliveryStatus contains a resolved URL to the dead letter sink address, and any other
resolved delivery options.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth provides the relevant information for OIDC authentication.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="eventing.knative.dev/v1alpha1">eventing.knative.dev/v1alpha1</h2>
<p>
<p>Package v1alpha1 is the v1alpha1 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#eventing.knative.dev/v1alpha1.EventPolicy">EventPolicy</a>
</li><li>
<a href="#eventing.knative.dev/v1alpha1.EventTransform">EventTransform</a>
</li><li>
<a href="#eventing.knative.dev/v1alpha1.RequestReply">RequestReply</a>
</li></ul>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicy">EventPolicy
</h3>
<p>
<p>EventPolicy represents a policy for addressable resources (Broker, Channel, sinks).</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>EventPolicy</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySpec">
EventPolicySpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventPolicy.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>to</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySpecTo">
[]EventPolicySpecTo
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>To lists all resources for which this policy applies.
Resources in this list must act like an ingress and have an audience.
The resources are part of the same namespace as the EventPolicy.
An empty list means it applies to all resources in the EventPolicies namespace</p>
</td>
</tr>
<tr>
<td>
<code>from</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySpecFrom">
[]EventPolicySpecFrom
</a>
</em>
</td>
<td>
<p>From is the list of sources or oidc identities, which are allowed to send events to the targets (.spec.to).</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is the list of SubscriptoinsApi filters which determine whether or not the event is accepted.
It is an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event will not
pass the target resource&rsquo;s ingress. Absence of any filters implies that the filters
always evaluate to true.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicyStatus">
EventPolicyStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventPolicy.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventTransform">EventTransform
</h3>
<p>
<p>EventTransform represents an even transformation.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>EventTransform</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventTransformSpec">
EventTransformSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventTransform.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sink is a reference to an object that will resolve to a uri to use as the sink.</p>
<p>If not present, the transformation will send back the transformed event as response, this
is useful to leverage the built-in Broker reply feature to re-publish a transformed event
back to the broker.</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.ReplySpec">
ReplySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is the configuration on how to handle responses from Sink.
It can only be set if Sink is set.</p>
</td>
</tr>
<tr>
<td>
<code>EventTransformations</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventTransformations">
EventTransformations
</a>
</em>
</td>
<td>
<p>
(Members of <code>EventTransformations</code> are embedded into this type.)
</p>
<p>EventTransformations contain all possible transformations, only one &ldquo;type&rdquo; can be used.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventTransformStatus">
EventTransformStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventTransform.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.RequestReply">RequestReply
</h3>
<p>
<p>RequestRepluy represents synchronous interface to sending and receiving events from a Broker.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>RequestReply</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.RequestReplySpec">
RequestReplySpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventPolicy.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>brokerRef</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>BrokerRef contains the reference to the broker the RequestReply sends events to.</p>
</td>
</tr>
<tr>
<td>
<code>correlationAttribute</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>replyAttribute</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>timeout</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.RequestReplyStatus">
RequestReplyStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventPolicy.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicyFromReference">EventPolicyFromReference
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicySpecFrom">EventPolicySpecFrom</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>API version of the referent.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind of the referent.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the referent.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names">https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</a></p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Namespace of the referent.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/">https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/</a>
This is optional field, it gets defaulted to the object holding it if left out.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicySelector">EventPolicySelector
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicySpecTo">EventPolicySpecTo</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>LabelSelector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<p>
(Members of <code>LabelSelector</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicySpec">EventPolicySpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicy">EventPolicy</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>to</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySpecTo">
[]EventPolicySpecTo
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>To lists all resources for which this policy applies.
Resources in this list must act like an ingress and have an audience.
The resources are part of the same namespace as the EventPolicy.
An empty list means it applies to all resources in the EventPolicies namespace</p>
</td>
</tr>
<tr>
<td>
<code>from</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySpecFrom">
[]EventPolicySpecFrom
</a>
</em>
</td>
<td>
<p>From is the list of sources or oidc identities, which are allowed to send events to the targets (.spec.to).</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is the list of SubscriptoinsApi filters which determine whether or not the event is accepted.
It is an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event will not
pass the target resource&rsquo;s ingress. Absence of any filters implies that the filters
always evaluate to true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicySpecFrom">EventPolicySpecFrom
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicySpec">EventPolicySpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicyFromReference">
EventPolicyFromReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ref contains a direct reference to a resource which is allowed to send events to the target.</p>
</td>
</tr>
<tr>
<td>
<code>sub</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sub sets the OIDC identity name to be allowed to send events to the target.
It is also possible to set a glob-like pattern to match any suffix.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicySpecTo">EventPolicySpecTo
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicySpec">EventPolicySpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ref</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicyToReference">
EventPolicyToReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ref contains the direct reference to a target</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventPolicySelector">
EventPolicySelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selector contains a selector to group targets</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicyStatus">EventPolicyStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicy">EventPolicy</a>)
</p>
<p>
<p>EventPolicyStatus represents the current state of a EventPolicy.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>from</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>From is the list of resolved oidc identities from .spec.from</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventPolicyToReference">EventPolicyToReference
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventPolicySpecTo">EventPolicySpecTo</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>API version of the referent.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind of the referent.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the referent.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names">https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventTransformSpec">EventTransformSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransform">EventTransform</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sink</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sink is a reference to an object that will resolve to a uri to use as the sink.</p>
<p>If not present, the transformation will send back the transformed event as response, this
is useful to leverage the built-in Broker reply feature to re-publish a transformed event
back to the broker.</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.ReplySpec">
ReplySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is the configuration on how to handle responses from Sink.
It can only be set if Sink is set.</p>
</td>
</tr>
<tr>
<td>
<code>EventTransformations</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventTransformations">
EventTransformations
</a>
</em>
</td>
<td>
<p>
(Members of <code>EventTransformations</code> are embedded into this type.)
</p>
<p>EventTransformations contain all possible transformations, only one &ldquo;type&rdquo; can be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventTransformStatus">EventTransformStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransform">EventTransform</a>)
</p>
<p>
<p>EventTransformStatus represents the current state of a EventTransform.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>SourceStatus inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the EventTransform fulfills the Addressable contract.
It exposes the endpoint as an URI to get events delivered.</p>
</td>
</tr>
<tr>
<td>
<code>jsonata</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.JsonataEventTransformationStatus">
JsonataEventTransformationStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>JsonataTransformationStatus is the status associated with JsonataEventTransformationSpec.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.EventTransformations">EventTransformations
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransformSpec">EventTransformSpec</a>, <a href="#eventing.knative.dev/v1alpha1.ReplySpec">ReplySpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>jsonata</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.JsonataEventTransformationSpec">
JsonataEventTransformationSpec
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.JsonataEventTransformationSpec">JsonataEventTransformationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransformations">EventTransformations</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>expression</code><br/>
<em>
string
</em>
</td>
<td>
<p>Expression is the JSONata expression.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.JsonataEventTransformationStatus">JsonataEventTransformationStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransformStatus">EventTransformStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>deployment</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#deploymentstatus-v1-apps">
Kubernetes apps/v1.DeploymentStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.ReplySpec">ReplySpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.EventTransformSpec">EventTransformSpec</a>)
</p>
<p>
<p>ReplySpec is the configurations on how to handle responses from Sink.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>EventTransformations</code><br/>
<em>
<a href="#eventing.knative.dev/v1alpha1.EventTransformations">
EventTransformations
</a>
</em>
</td>
<td>
<p>
(Members of <code>EventTransformations</code> are embedded into this type.)
</p>
<p>EventTransformations for replies from the Sink, contain all possible transformations,
only one &ldquo;type&rdquo; can be used.</p>
<p>The used type must match the top-level transformation, if you need to mix transformation types,
use compositions and chain transformations together to achieve your desired outcome.</p>
</td>
</tr>
<tr>
<td>
<code>discard</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Discard discards responses from Sink and return empty response body.</p>
<p>When set to false, it returns the exact sink response body.
When set to true, Discard is mutually exclusive with EventTransformations in the reply
section, it can either be discarded or transformed.</p>
<p>Default: false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.RequestReplySpec">RequestReplySpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.RequestReply">RequestReply</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>brokerRef</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>BrokerRef contains the reference to the broker the RequestReply sends events to.</p>
</td>
</tr>
<tr>
<td>
<code>correlationAttribute</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>replyAttribute</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>timeout</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1alpha1.RequestReplyStatus">RequestReplyStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1alpha1.RequestReply">RequestReply</a>)
</p>
<p>
<p>RequestReplyStatus represents the current state of a RequestReply.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the RequestReply fulfills the Addressable contract.
It exposes the endpoint as an URI to get events delivered.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this Broker.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="eventing.knative.dev/v1beta1">eventing.knative.dev/v1beta1</h2>
<p>
<p>Package v1beta1 is the v1beta1 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#eventing.knative.dev/v1beta1.EventType">EventType</a>
</li></ul>
<h3 id="eventing.knative.dev/v1beta1.EventType">EventType
</h3>
<p>
<p>EventType represents a type of event that can be consumed from a Broker.
Deprecated: use v1beta2.EventType instead.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1beta1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>EventType</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta1.EventTypeSpec">
EventTypeSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventType.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type represents the CloudEvents type. It is authoritative.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source is a URI, it represents the CloudEvents source.</p>
</td>
</tr>
<tr>
<td>
<code>schema</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schema is a URI, it represents the CloudEvents schemaurl extension attribute.
It may be a JSON schema, a protobuf schema, etc. It is optional.</p>
</td>
</tr>
<tr>
<td>
<code>schemaData</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchemaData allows the CloudEvents schema to be stored directly in the
EventType. Content is dependent on the encoding. Optional attribute.
The contents are not validated or manipulated by the system.</p>
</td>
</tr>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Broker refers to the Broker that can provide the EventType.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta1.EventTypeStatus">
EventTypeStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventType.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta1.EventTypeSpec">EventTypeSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta1.EventType">EventType</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type represents the CloudEvents type. It is authoritative.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source is a URI, it represents the CloudEvents source.</p>
</td>
</tr>
<tr>
<td>
<code>schema</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schema is a URI, it represents the CloudEvents schemaurl extension attribute.
It may be a JSON schema, a protobuf schema, etc. It is optional.</p>
</td>
</tr>
<tr>
<td>
<code>schemaData</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchemaData allows the CloudEvents schema to be stored directly in the
EventType. Content is dependent on the encoding. Optional attribute.
The contents are not validated or manipulated by the system.</p>
</td>
</tr>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Broker refers to the Broker that can provide the EventType.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta1.EventTypeStatus">EventTypeStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta1.EventType">EventType</a>)
</p>
<p>
<p>EventTypeStatus represents the current state of a EventType.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="eventing.knative.dev/v1beta2">eventing.knative.dev/v1beta2</h2>
<p>
<p>Package v1beta2 is the v1beta2 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#eventing.knative.dev/v1beta2.EventType">EventType</a>
</li></ul>
<h3 id="eventing.knative.dev/v1beta2.EventType">EventType
</h3>
<p>
<p>EventType represents a type of event that can be consumed from a Broker.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1beta2
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>EventType</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta2.EventTypeSpec">
EventTypeSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventType.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type represents the CloudEvents type. It is authoritative.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source is a URI, it represents the CloudEvents source.</p>
</td>
</tr>
<tr>
<td>
<code>schema</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schema is a URI, it represents the CloudEvents schemaurl extension attribute.
It may be a JSON schema, a protobuf schema, etc. It is optional.</p>
</td>
</tr>
<tr>
<td>
<code>schemaData</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchemaData allows the CloudEvents schema to be stored directly in the
EventType. Content is dependent on the encoding. Optional attribute.
The contents are not validated or manipulated by the system.</p>
</td>
</tr>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Broker refers to the Broker that can provide the EventType.
Deprecated: This field is deprecated and will be removed in a future release.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta2.EventTypeStatus">
EventTypeStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventType.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta2.EventTypeSpec">EventTypeSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta2.EventType">EventType</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type represents the CloudEvents type. It is authoritative.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source is a URI, it represents the CloudEvents source.</p>
</td>
</tr>
<tr>
<td>
<code>schema</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schema is a URI, it represents the CloudEvents schemaurl extension attribute.
It may be a JSON schema, a protobuf schema, etc. It is optional.</p>
</td>
</tr>
<tr>
<td>
<code>schemaData</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchemaData allows the CloudEvents schema to be stored directly in the
EventType. Content is dependent on the encoding. Optional attribute.
The contents are not validated or manipulated by the system.</p>
</td>
</tr>
<tr>
<td>
<code>broker</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Broker refers to the Broker that can provide the EventType.
Deprecated: This field is deprecated and will be removed in a future release.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta2.EventTypeStatus">EventTypeStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta2.EventType">EventType</a>)
</p>
<p>
<p>EventTypeStatus represents the current state of a EventType.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="eventing.knative.dev/v1beta3">eventing.knative.dev/v1beta3</h2>
<p>
<p>Package v1beta3 is the v1beta3 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#eventing.knative.dev/v1beta3.EventType">EventType</a>
</li></ul>
<h3 id="eventing.knative.dev/v1beta3.EventType">EventType
</h3>
<p>
<p>EventType represents a type of event that can be consumed from a Broker.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
eventing.knative.dev/v1beta3
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>EventType</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta3.EventTypeSpec">
EventTypeSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the EventType.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
<tr>
<td>
<code>attributes</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta3.EventAttributeDefinition">
[]EventAttributeDefinition
</a>
</em>
</td>
<td>
<p>Attributes is an array of CloudEvent attributes and extension attributes.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta3.EventTypeStatus">
EventTypeStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the EventType.
This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta3.EventAttributeDefinition">EventAttributeDefinition
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta3.EventTypeSpec">EventTypeSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the CloudEvents attribute.</p>
</td>
</tr>
<tr>
<td>
<code>required</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Required determines whether this attribute must be set on corresponding CloudEvents.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value is a string representing the allowable values for the EventType attribute.
It may be a single value such as &ldquo;/apis/v1/namespaces/default/pingsource/ps&rdquo;, or it could be a template
for the allowed values, such as &ldquo;/apis/v1/namespaces/{namespace}/pingsource/{sourceName}.
To specify a section of the string value which may change between different CloudEvents
you can use curly brackets {} and optionally a variable name between them.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta3.EventTypeSpec">EventTypeSpec
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta3.EventType">EventType</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference is a KReference to the belonging addressable.
For example, this could be a pointer to a Broker.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description is an optional field used to describe the EventType, in any meaningful way.</p>
</td>
</tr>
<tr>
<td>
<code>attributes</code><br/>
<em>
<a href="#eventing.knative.dev/v1beta3.EventAttributeDefinition">
[]EventAttributeDefinition
</a>
</em>
</td>
<td>
<p>Attributes is an array of CloudEvent attributes and extension attributes.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventing.knative.dev/v1beta3.EventTypeStatus">EventTypeStatus
</h3>
<p>
(<em>Appears on:</em><a href="#eventing.knative.dev/v1beta3.EventType">EventType</a>)
</p>
<p>
<p>EventTypeStatus represents the current state of a EventType.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="flows.knative.dev/v1">flows.knative.dev/v1</h2>
<p>
<p>Package v1 is the v1 version of the API.</p>
</p>
Resource Types:
<ul></ul>
<h3 id="flows.knative.dev/v1.Parallel">Parallel
</h3>
<p>
<p>Parallel defines conditional branches that will be wired in
series through Channels and Subscriptions.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelSpec">
ParallelSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Parallel.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>branches</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelBranch">
[]ParallelBranch
</a>
</em>
</td>
<td>
<p>Branches is the list of Filter/Subscribers pairs.</p>
</td>
</tr>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use. If left unspecified, it is set to the default Channel CRD
for the namespace (or cluster, in case there are no defaults for the namespace).</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is a Reference to where the result of a case Subscriber gets sent to
when the case does not have a Reply</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelStatus">
ParallelStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Parallel. This data may be out of
date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelBranch">ParallelBranch
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.ParallelSpec">ParallelSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>filter</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filter is the expression guarding the branch</p>
</td>
</tr>
<tr>
<td>
<code>subscriber</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>Subscriber receiving the event when the filter passes</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is a Reference to where the result of Subscriber of this case gets sent to.
If not specified, sent the result to the Parallel Reply</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery is the delivery specification for events to the subscriber
This includes things like retries, DLS, etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelBranchStatus">ParallelBranchStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.ParallelStatus">ParallelStatus</a>)
</p>
<p>
<p>ParallelBranchStatus represents the current state of a Parallel branch</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>filterSubscriptionStatus</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelSubscriptionStatus">
ParallelSubscriptionStatus
</a>
</em>
</td>
<td>
<p>FilterSubscriptionStatus corresponds to the filter subscription status.</p>
</td>
</tr>
<tr>
<td>
<code>filterChannelStatus</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelChannelStatus">
ParallelChannelStatus
</a>
</em>
</td>
<td>
<p>FilterChannelStatus corresponds to the filter channel status.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberSubscriptionStatus</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelSubscriptionStatus">
ParallelSubscriptionStatus
</a>
</em>
</td>
<td>
<p>SubscriptionStatus corresponds to the subscriber subscription status.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelChannelStatus">ParallelChannelStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.ParallelBranchStatus">ParallelBranchStatus</a>, <a href="#flows.knative.dev/v1.ParallelStatus">ParallelStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>channel</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>Channel is the reference to the underlying channel.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#Condition">
knative.dev/pkg/apis.Condition
</a>
</em>
</td>
<td>
<p>ReadyCondition indicates whether the Channel is ready or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelSpec">ParallelSpec
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.Parallel">Parallel</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>branches</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelBranch">
[]ParallelBranch
</a>
</em>
</td>
<td>
<p>Branches is the list of Filter/Subscribers pairs.</p>
</td>
</tr>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use. If left unspecified, it is set to the default Channel CRD
for the namespace (or cluster, in case there are no defaults for the namespace).</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is a Reference to where the result of a case Subscriber gets sent to
when the case does not have a Reply</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelStatus">ParallelStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.Parallel">Parallel</a>)
</p>
<p>
<p>ParallelStatus represents the current state of a Parallel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>ingressChannelStatus</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelChannelStatus">
ParallelChannelStatus
</a>
</em>
</td>
<td>
<p>IngressChannelStatus corresponds to the ingress channel status.</p>
</td>
</tr>
<tr>
<td>
<code>branchStatuses</code><br/>
<em>
<a href="#flows.knative.dev/v1.ParallelBranchStatus">
[]ParallelBranchStatus
</a>
</em>
</td>
<td>
<p>BranchStatuses is an array of corresponding to branch statuses.
Matches the Spec.Branches array in the order.</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<p>AddressStatus is the starting point to this Parallel. Sending to this
will target the first subscriber.
It generally has the form {channel}.{namespace}.svc.{cluster domain name}</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth provides the relevant information for OIDC authentication.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this Broker</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.ParallelSubscriptionStatus">ParallelSubscriptionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.ParallelBranchStatus">ParallelBranchStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscription</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>Subscription is the reference to the underlying Subscription.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#Condition">
knative.dev/pkg/apis.Condition
</a>
</em>
</td>
<td>
<p>ReadyCondition indicates whether the Subscription is ready or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.Sequence">Sequence
</h3>
<p>
<p>Sequence defines a sequence of Subscribers that will be wired in
series through Channels and Subscriptions.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceSpec">
SequenceSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Sequence.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>steps</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceStep">
[]SequenceStep
</a>
</em>
</td>
<td>
<p>Steps is the list of Destinations (processors / functions) that will be called in the order
provided. Each step has its own delivery options</p>
</td>
</tr>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use. If left unspecified, it is set to the default Channel CRD
for the namespace (or cluster, in case there are no defaults for the namespace).</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is a Reference to where the result of the last Subscriber gets sent to.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceStatus">
SequenceStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Sequence. This data may be out of
date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.SequenceChannelStatus">SequenceChannelStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.SequenceStatus">SequenceStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>channel</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>Channel is the reference to the underlying channel.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#Condition">
knative.dev/pkg/apis.Condition
</a>
</em>
</td>
<td>
<p>ReadyCondition indicates whether the Channel is ready or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.SequenceSpec">SequenceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.Sequence">Sequence</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>steps</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceStep">
[]SequenceStep
</a>
</em>
</td>
<td>
<p>Steps is the list of Destinations (processors / functions) that will be called in the order
provided. Each step has its own delivery options</p>
</td>
</tr>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use. If left unspecified, it is set to the default Channel CRD
for the namespace (or cluster, in case there are no defaults for the namespace).</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply is a Reference to where the result of the last Subscriber gets sent to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.SequenceStatus">SequenceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.Sequence">Sequence</a>)
</p>
<p>
<p>SequenceStatus represents the current state of a Sequence.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionStatuses</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceSubscriptionStatus">
[]SequenceSubscriptionStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriptionStatuses is an array of corresponding Subscription statuses.
Matches the Spec.Steps array in the order.</p>
</td>
</tr>
<tr>
<td>
<code>channelStatuses</code><br/>
<em>
<a href="#flows.knative.dev/v1.SequenceChannelStatus">
[]SequenceChannelStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelStatuses is an array of corresponding Channel statuses.
Matches the Spec.Steps array in the order.</p>
</td>
</tr>
<tr>
<td>
<code>address</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Addressable">
knative.dev/pkg/apis/duck/v1.Addressable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Address is the starting point to this Sequence. Sending to this
will target the first subscriber.
It generally has the form {channel}.{namespace}.svc.{cluster domain name}</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth provides the relevant information for OIDC authentication.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this Broker</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.SequenceStep">SequenceStep
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.SequenceSpec">SequenceSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Destination</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>
(Members of <code>Destination</code> are embedded into this type.)
</p>
<p>Subscriber receiving the step event</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery is the delivery specification for events to the subscriber
This includes things like retries, DLS, etc.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="flows.knative.dev/v1.SequenceSubscriptionStatus">SequenceSubscriptionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.SequenceStatus">SequenceStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscription</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectreference-v1-core">
Kubernetes core/v1.ObjectReference
</a>
</em>
</td>
<td>
<p>Subscription is the reference to the underlying Subscription.</p>
</td>
</tr>
<tr>
<td>
<code>ready</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#Condition">
knative.dev/pkg/apis.Condition
</a>
</em>
</td>
<td>
<p>ReadyCondition indicates whether the Subscription is ready or not.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="messaging.knative.dev/v1">messaging.knative.dev/v1</h2>
<p>
<p>Package v1 is the v1 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#messaging.knative.dev/v1.Channel">Channel</a>
</li><li>
<a href="#messaging.knative.dev/v1.InMemoryChannel">InMemoryChannel</a>
</li><li>
<a href="#messaging.knative.dev/v1.Subscription">Subscription</a>
</li></ul>
<h3 id="messaging.knative.dev/v1.Channel">Channel
</h3>
<p>
<p>Channel represents a generic Channel. It is normally used when we want a
Channel, but do not need a specific Channel implementation.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
messaging.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Channel</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelSpec">
ChannelSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Channel.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use to create the CRD
Channel backing this Channel. This is immutable after creation.
Normally this is set by the Channel defaulter, not directly by the user.</p>
</td>
</tr>
<tr>
<td>
<code>ChannelableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableSpec</code> are embedded into this type.)
</p>
<p>Channel conforms to ChannelableSpec</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelStatus">
ChannelStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Channel. This data may be out
of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.InMemoryChannel">InMemoryChannel
</h3>
<p>
<p>InMemoryChannel is a resource representing an in memory channel</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
messaging.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>InMemoryChannel</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#messaging.knative.dev/v1.InMemoryChannelSpec">
InMemoryChannelSpec
</a>
</em>
</td>
<td>
<p>Spec defines the desired state of the Channel.</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>ChannelableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableSpec</code> are embedded into this type.)
</p>
<p>Channel conforms to Duck type Channelable.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#messaging.knative.dev/v1.InMemoryChannelStatus">
InMemoryChannelStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Status represents the current state of the Channel. This data may be out of
date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.Subscription">Subscription
</h3>
<p>
<p>Subscription routes events received on a Channel to a DNS name and
corresponds to the subscriptions.channels.knative.dev CRD.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
messaging.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Subscription</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#messaging.knative.dev/v1.SubscriptionSpec">
SubscriptionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>channel</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>Reference to a channel that will be used to create the subscription
You can specify only the following fields of the KReference:
- Kind
- APIVersion
- Name
- Namespace
The resource pointed by this KReference must meet the
contract to the ChannelableSpec duck type. If the resource does not
meet this contract it will be reflected in the Subscription&rsquo;s status.</p>
<p>This field is immutable. We have no good answer on what happens to
the events that are currently in the channel being consumed from
and what the semantics there should be. For now, you can always
delete the Subscription and recreate it to point to a different
channel, giving the user more control over what semantics should
be used (drain the channel first, possibly have events dropped,
etc.)</p>
</td>
</tr>
<tr>
<td>
<code>subscriber</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>Subscriber is reference to function for processing events.
Events from the Channel will be delivered here and replies are
sent to a Destination as specified by the Reply.</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply specifies (optionally) how to handle events returned from
the Subscriber target.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery configuration</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#messaging.knative.dev/v1.SubscriptionStatus">
SubscriptionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.ChannelDefaulter">ChannelDefaulter
</h3>
<p>
<p>ChannelDefaulter sets the default Channel CRD and Arguments on Channels that do not
specify any implementation.</p>
</p>
<h3 id="messaging.knative.dev/v1.ChannelSpec">ChannelSpec
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.Channel">Channel</a>)
</p>
<p>
<p>ChannelSpec defines which subscribers have expressed interest in receiving
events from this Channel. It also defines the ChannelTemplate to use in
order to create the CRD Channel backing this Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>channelTemplate</code><br/>
<em>
<a href="#messaging.knative.dev/v1.ChannelTemplateSpec">
ChannelTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ChannelTemplate specifies which Channel CRD to use to create the CRD
Channel backing this Channel. This is immutable after creation.
Normally this is set by the Channel defaulter, not directly by the user.</p>
</td>
</tr>
<tr>
<td>
<code>ChannelableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableSpec</code> are embedded into this type.)
</p>
<p>Channel conforms to ChannelableSpec</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.ChannelStatus">ChannelStatus
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.Channel">Channel</a>)
</p>
<p>
<p>ChannelStatus represents the current state of a Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ChannelableStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableStatus">
ChannelableStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableStatus</code> are embedded into this type.)
</p>
<p>Channel conforms to ChannelableStatus</p>
</td>
</tr>
<tr>
<td>
<code>channel</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Channel is an KReference to the Channel CRD backing this Channel.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.ChannelTemplateSpec">ChannelTemplateSpec
</h3>
<p>
(<em>Appears on:</em><a href="#flows.knative.dev/v1.ParallelSpec">ParallelSpec</a>, <a href="#flows.knative.dev/v1.SequenceSpec">SequenceSpec</a>, <a href="#messaging.knative.dev/v1.ChannelSpec">ChannelSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>spec</code><br/>
<em>
k8s.io/apimachinery/pkg/runtime.RawExtension
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec defines the Spec to use for each channel created. Passed
in verbatim to the Channel CRD as Spec section.</p>
<br/>
<br/>
<table>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.ChannelTemplateSpecOption">ChannelTemplateSpecOption
</h3>
<p>
<p>ChannelTemplateSpecOption is an optional function for ChannelTemplateSpec.</p>
</p>
<h3 id="messaging.knative.dev/v1.InMemoryChannelSpec">InMemoryChannelSpec
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.InMemoryChannel">InMemoryChannel</a>)
</p>
<p>
<p>InMemoryChannelSpec defines which subscribers have expressed interest in
receiving events from this InMemoryChannel.
arguments for a Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ChannelableSpec</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableSpec">
ChannelableSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableSpec</code> are embedded into this type.)
</p>
<p>Channel conforms to Duck type Channelable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.InMemoryChannelStatus">InMemoryChannelStatus
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.InMemoryChannel">InMemoryChannel</a>)
</p>
<p>
<p>ChannelStatus represents the current state of a Channel.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ChannelableStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.ChannelableStatus">
ChannelableStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>ChannelableStatus</code> are embedded into this type.)
</p>
<p>Channel conforms to Duck type ChannelableStatus.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.SubscriptionSpec">SubscriptionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.Subscription">Subscription</a>)
</p>
<p>
<p>SubscriptionSpec specifies the Channel for incoming events, a Subscriber target
for processing those events and where to put the result of the processing. Only
From (where the events are coming from) is always required. You can optionally
only Process the events (results in no output events) by leaving out the Reply.
You can also perform an identity transformation on the incoming events by leaving
out the Subscriber and only specifying Reply.</p>
<p>The following are all valid specifications:
channel &ndash;[subscriber]&ndash;&gt; reply
Sink, no outgoing events:
channel &ndash; subscriber
no-op function (identity transformation):
channel &ndash;&gt; reply</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>channel</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#KReference">
knative.dev/pkg/apis/duck/v1.KReference
</a>
</em>
</td>
<td>
<p>Reference to a channel that will be used to create the subscription
You can specify only the following fields of the KReference:
- Kind
- APIVersion
- Name
- Namespace
The resource pointed by this KReference must meet the
contract to the ChannelableSpec duck type. If the resource does not
meet this contract it will be reflected in the Subscription&rsquo;s status.</p>
<p>This field is immutable. We have no good answer on what happens to
the events that are currently in the channel being consumed from
and what the semantics there should be. For now, you can always
delete the Subscription and recreate it to point to a different
channel, giving the user more control over what semantics should
be used (drain the channel first, possibly have events dropped,
etc.)</p>
</td>
</tr>
<tr>
<td>
<code>subscriber</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<p>Subscriber is reference to function for processing events.
Events from the Channel will be delivered here and replies are
sent to a Destination as specified by the Reply.</p>
</td>
</tr>
<tr>
<td>
<code>reply</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Destination">
knative.dev/pkg/apis/duck/v1.Destination
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reply specifies (optionally) how to handle events returned from
the Subscriber target.</p>
</td>
</tr>
<tr>
<td>
<code>delivery</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliverySpec">
DeliverySpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Delivery configuration</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.SubscriptionStatus">SubscriptionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.Subscription">Subscription</a>)
</p>
<p>
<p>SubscriptionStatus (computed) for a subscription</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
<p>inherits duck/v1 Status, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current state.</p>
</td>
</tr>
<tr>
<td>
<code>physicalSubscription</code><br/>
<em>
<a href="#messaging.knative.dev/v1.SubscriptionStatusPhysicalSubscription">
SubscriptionStatusPhysicalSubscription
</a>
</em>
</td>
<td>
<p>PhysicalSubscription is the fully resolved values that this Subscription represents.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AuthStatus">
knative.dev/pkg/apis/duck/v1.AuthStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Auth provides the relevant information for OIDC authentication.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="messaging.knative.dev/v1.SubscriptionStatusPhysicalSubscription">SubscriptionStatusPhysicalSubscription
</h3>
<p>
(<em>Appears on:</em><a href="#messaging.knative.dev/v1.SubscriptionStatus">SubscriptionStatus</a>)
</p>
<p>
<p>SubscriptionStatusPhysicalSubscription represents the fully resolved values for this
Subscription.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subscriberUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberURI is the fully resolved URI for spec.subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberCACerts is the Certification Authority (CA) certificates in PEM
format according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a> for the
resolved URI for spec.subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>subscriberAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubscriberAudience is the OIDC audience for the the resolved URI for
spec.subscriber.</p>
</td>
</tr>
<tr>
<td>
<code>replyUri</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis#URL">
knative.dev/pkg/apis.URL
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyURI is the fully resolved URI for the spec.reply.</p>
</td>
</tr>
<tr>
<td>
<code>replyCACerts</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyCACerts is the Certification Authority (CA) certificates in PEM
format according to <a href="https://www.rfc-editor.org/rfc/rfc7468">https://www.rfc-editor.org/rfc/rfc7468</a> for the
resolved URI for the spec.reply.</p>
</td>
</tr>
<tr>
<td>
<code>replyAudience</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReplyAudience is the OIDC audience for the the resolved URI for
spec.reply.</p>
</td>
</tr>
<tr>
<td>
<code>DeliveryStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.DeliveryStatus">
DeliveryStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>DeliveryStatus</code> are embedded into this type.)
</p>
<p>DeliveryStatus contains a resolved URL to the dead letter sink address, and any other
resolved delivery options.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="sinks.knative.dev/v1alpha1">sinks.knative.dev/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains API Schema definitions for the sources v1alpha1 API group.</p>
</p>
Resource Types:
<ul><li>
<a href="#sinks.knative.dev/v1alpha1.IntegrationSink">IntegrationSink</a>
</li><li>
<a href="#sinks.knative.dev/v1alpha1.JobSink">JobSink</a>
</li></ul>
<h3 id="sinks.knative.dev/v1alpha1.IntegrationSink">IntegrationSink
</h3>
<p>
<p>IntegrationSink is the Schema for the IntegrationSink API.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sinks.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>IntegrationSink</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.IntegrationSinkSpec">
IntegrationSinkSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>aws</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.Aws">
Aws
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>log</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.Log">
Log
</a>
</em>
</td>
<td>
<p>AWS source configuration</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.IntegrationSinkStatus">
IntegrationSinkStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.JobSink">JobSink
</h3>
<p>
<p>JobSink is the Schema for the JobSink API.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sinks.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>JobSink</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.JobSinkSpec">
JobSinkSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>job</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#job-v1-batch">
Kubernetes batch/v1.Job
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Job to run when an event occur.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.JobSinkStatus">
JobSinkStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.Aws">Aws
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.IntegrationSinkSpec">IntegrationSinkSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>s3</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSS3
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sqs</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSSQS
</em>
</td>
<td>
<p>S3 source configuration</p>
</td>
</tr>
<tr>
<td>
<code>sns</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSSNS
</em>
</td>
<td>
<p>SQS source configuration</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.Auth
</em>
</td>
<td>
<p>SNS source configuration</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.ExecutionMode">ExecutionMode
(<code>string</code> alias)</p></h3>
<p>
</p>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;batch&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.IntegrationSinkSpec">IntegrationSinkSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.IntegrationSink">IntegrationSink</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>aws</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.Aws">
Aws
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>log</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.Log">
Log
</a>
</em>
</td>
<td>
<p>AWS source configuration</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.IntegrationSinkStatus">IntegrationSinkStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.IntegrationSink">IntegrationSink</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the JobSink fulfills the Addressable contract.
It exposes the endpoint as an URI to get events delivered.</p>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this JobSink</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.JobSinkSpec">JobSinkSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.JobSink">JobSink</a>)
</p>
<p>
<p>JobSinkSpec defines the desired state of the JobSink.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>job</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#job-v1-batch">
Kubernetes batch/v1.Job
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Job to run when an event occur.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.JobSinkStatus">JobSinkStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.JobSink">JobSink</a>)
</p>
<p>
<p>JobSinkStatus defines the observed state of JobSink.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Status</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#Status">
knative.dev/pkg/apis/duck/v1.Status
</a>
</em>
</td>
<td>
<p>
(Members of <code>Status</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>AddressStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#AddressStatus">
knative.dev/pkg/apis/duck/v1.AddressStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AddressStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AddressStatus is the part where the JobSink fulfills the Addressable contract.
It exposes the endpoint as an URI to get events delivered.</p>
</td>
</tr>
<tr>
<td>
<code>job</code><br/>
<em>
<a href="#sinks.knative.dev/v1alpha1.JobStatus">
JobStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>AppliedEventPoliciesStatus</code><br/>
<em>
<a href="#duck.knative.dev/v1.AppliedEventPoliciesStatus">
AppliedEventPoliciesStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>AppliedEventPoliciesStatus</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>AppliedEventPoliciesStatus contains the list of EventPolicies which apply to this JobSink</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.JobStatus">JobStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.JobSinkStatus">JobSinkStatus</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>selector</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sinks.knative.dev/v1alpha1.Log">Log
</h3>
<p>
(<em>Appears on:</em><a href="#sinks.knative.dev/v1alpha1.IntegrationSinkSpec">IntegrationSinkSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>loggerName</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the logging category to use</p>
</td>
</tr>
<tr>
<td>
<code>logMask</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Logging level to use</p>
</td>
</tr>
<tr>
<td>
<code>marker</code><br/>
<em>
string
</em>
</td>
<td>
<p>Mask sensitive information in the log</p>
</td>
</tr>
<tr>
<td>
<code>multiline</code><br/>
<em>
bool
</em>
</td>
<td>
<p>An optional Marker name to use</p>
</td>
</tr>
<tr>
<td>
<code>showAllProperties</code><br/>
<em>
bool
</em>
</td>
<td>
<p>If enabled, outputs each information on a newline</p>
</td>
</tr>
<tr>
<td>
<code>showBody</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show all of the exchange properties (both internal and custom)</p>
</td>
</tr>
<tr>
<td>
<code>showBodyType</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the message body</p>
</td>
</tr>
<tr>
<td>
<code>showExchangePattern</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the body Java type</p>
</td>
</tr>
<tr>
<td>
<code>showHeaders</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the Message Exchange Pattern (MEP)</p>
</td>
</tr>
<tr>
<td>
<code>showProperties</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the headers received</p>
</td>
</tr>
<tr>
<td>
<code>showStreams</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the exchange properties (only custom)</p>
</td>
</tr>
<tr>
<td>
<code>showCachedStreams</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Show the stream bodies</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="sources.knative.dev/v1">sources.knative.dev/v1</h2>
<p>
<p>Package v1 contains API Schema definitions for the sources v1 API group.</p>
</p>
Resource Types:
<ul><li>
<a href="#sources.knative.dev/v1.ApiServerSource">ApiServerSource</a>
</li><li>
<a href="#sources.knative.dev/v1.ContainerSource">ContainerSource</a>
</li><li>
<a href="#sources.knative.dev/v1.PingSource">PingSource</a>
</li><li>
<a href="#sources.knative.dev/v1.SinkBinding">SinkBinding</a>
</li></ul>
<h3 id="sources.knative.dev/v1.ApiServerSource">ApiServerSource
</h3>
<p>
<p>ApiServerSource is the Schema for the apiserversources API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ApiServerSource</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1.ApiServerSourceSpec">
ApiServerSourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#sources.knative.dev/v1.APIVersionKindSelector">
[]APIVersionKindSelector
</a>
</em>
</td>
<td>
<p>Resource are the resources this source will track and send related
lifecycle events from the Kubernetes ApiServer, with an optional label
selector to help filter.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="#sources.knative.dev/v1.APIVersionKind">
APIVersionKind
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResourceOwner is an additional filter to only track resources that are
owned by a specific resource type. If ResourceOwner matches Resources[n]
then Resources[n] is allowed to pass the ResourceOwner filter.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>EventMode controls the format of the event.
<code>Reference</code> sends a dataref event type for the resource under watch.
<code>Resource</code> send the full resource lifecycle event.
Defaults to <code>Reference</code></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceAccountName is the name of the ServiceAccount to use to run this
source. Defaults to default if not set.</p>
</td>
</tr>
<tr>
<td>
<code>namespaceSelector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>NamespaceSelector is a label selector to capture the namespaces that
should be watched by the source.</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is an experimental field that conforms to the CNCF CloudEvents Subscriptions
API. It&rsquo;s an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event MUST
NOT be sent to the Sink. If all the filter expressions in the array
evaluate to true, the event MUST be attempted to be delivered. Absence of
a filter or empty array implies a value of true.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1.ApiServerSourceStatus">
ApiServerSourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.ContainerSource">ContainerSource
</h3>
<p>
<p>ContainerSource is the Schema for the containersources API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ContainerSource</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1.ContainerSourceSpec">
ContainerSourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
<p>Template describes the pods that will be created</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1.ContainerSourceStatus">
ContainerSourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.PingSource">PingSource
</h3>
<p>
<p>PingSource is the Schema for the PingSources API.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>PingSource</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1.PingSourceSpec">
PingSourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>schedule</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schedule is the cron schedule. Defaults to <code>* * * * *</code>.</p>
</td>
</tr>
<tr>
<td>
<code>timezone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timezone modifies the actual time relative to the specified timezone.
Defaults to the system time zone.
More general information about time zones: <a href="https://www.iana.org/time-zones">https://www.iana.org/time-zones</a>
List of valid timezone values: <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">https://en.wikipedia.org/wiki/List_of_tz_database_time_zones</a></p>
</td>
</tr>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContentType is the media type of Data or DataBase64. Default is empty.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is data used as the body of the event posted to the sink. Default is empty.
Mutually exclusive with DataBase64.</p>
</td>
</tr>
<tr>
<td>
<code>dataBase64</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DataBase64 is the base64-encoded string of the actual event&rsquo;s body posted to the sink. Default is empty.
Mutually exclusive with Data.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1.PingSourceStatus">
PingSourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.SinkBinding">SinkBinding
</h3>
<p>
<p>SinkBinding describes a Binding that is also a Source.
The <code>sink</code> (from the Source duck) is resolved to a URL and
then projected into the <code>subject</code> by augmenting the runtime
contract of the referenced containers to have a <code>K_SINK</code>
environment variable holding the endpoint to which to send
cloud events.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>SinkBinding</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1.SinkBindingSpec">
SinkBindingSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>BindingSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#BindingSpec">
knative.dev/pkg/apis/duck/v1.BindingSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>BindingSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 BindingSpec, which currently provides:
* Subject - Subject references the resource(s) whose &ldquo;runtime contract&rdquo;
should be augmented by Binding implementations.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1.SinkBindingStatus">
SinkBindingStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.APIVersionKind">APIVersionKind
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ApiServerSourceSpec">ApiServerSourceSpec</a>)
</p>
<p>
<p>APIVersionKind is an APIVersion and Kind tuple.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion - the API version of the resource to watch.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind of the resource to watch.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.APIVersionKindSelector">APIVersionKindSelector
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ApiServerSourceSpec">ApiServerSourceSpec</a>)
</p>
<p>
<p>APIVersionKindSelector is an APIVersion Kind tuple with a LabelSelector.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion - the API version of the resource to watch.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind of the resource to watch.
More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LabelSelector filters this source to objects to those resources pass the
label selector.
More info: <a href="http://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors">http://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.ApiServerSourceSpec">ApiServerSourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ApiServerSource">ApiServerSource</a>)
</p>
<p>
<p>ApiServerSourceSpec defines the desired state of ApiServerSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#sources.knative.dev/v1.APIVersionKindSelector">
[]APIVersionKindSelector
</a>
</em>
</td>
<td>
<p>Resource are the resources this source will track and send related
lifecycle events from the Kubernetes ApiServer, with an optional label
selector to help filter.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="#sources.knative.dev/v1.APIVersionKind">
APIVersionKind
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ResourceOwner is an additional filter to only track resources that are
owned by a specific resource type. If ResourceOwner matches Resources[n]
then Resources[n] is allowed to pass the ResourceOwner filter.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>EventMode controls the format of the event.
<code>Reference</code> sends a dataref event type for the resource under watch.
<code>Resource</code> send the full resource lifecycle event.
Defaults to <code>Reference</code></p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceAccountName is the name of the ServiceAccount to use to run this
source. Defaults to default if not set.</p>
</td>
</tr>
<tr>
<td>
<code>namespaceSelector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>NamespaceSelector is a label selector to capture the namespaces that
should be watched by the source.</p>
</td>
</tr>
<tr>
<td>
<code>filters</code><br/>
<em>
<a href="#eventing.knative.dev/v1.SubscriptionsAPIFilter">
[]SubscriptionsAPIFilter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Filters is an experimental field that conforms to the CNCF CloudEvents Subscriptions
API. It&rsquo;s an array of filter expressions that evaluate to true or false.
If any filter expression in the array evaluates to false, the event MUST
NOT be sent to the Sink. If all the filter expressions in the array
evaluate to true, the event MUST be attempted to be delivered. Absence of
a filter or empty array implies a value of true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.ApiServerSourceStatus">ApiServerSourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ApiServerSource">ApiServerSource</a>)
</p>
<p>
<p>ApiServerSourceStatus defines the observed state of ApiServerSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
<tr>
<td>
<code>namespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Namespaces show the namespaces currently watched by the ApiServerSource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.ContainerSourceSpec">ContainerSourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ContainerSource">ContainerSource</a>)
</p>
<p>
<p>ContainerSourceSpec defines the desired state of ContainerSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
<p>Template describes the pods that will be created</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.ContainerSourceStatus">ContainerSourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.ContainerSource">ContainerSource</a>)
</p>
<p>
<p>ContainerSourceStatus defines the observed state of ContainerSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.PingSourceSpec">PingSourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.PingSource">PingSource</a>)
</p>
<p>
<p>PingSourceSpec defines the desired state of the PingSource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>schedule</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schedule is the cron schedule. Defaults to <code>* * * * *</code>.</p>
</td>
</tr>
<tr>
<td>
<code>timezone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timezone modifies the actual time relative to the specified timezone.
Defaults to the system time zone.
More general information about time zones: <a href="https://www.iana.org/time-zones">https://www.iana.org/time-zones</a>
List of valid timezone values: <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">https://en.wikipedia.org/wiki/List_of_tz_database_time_zones</a></p>
</td>
</tr>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContentType is the media type of Data or DataBase64. Default is empty.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is data used as the body of the event posted to the sink. Default is empty.
Mutually exclusive with DataBase64.</p>
</td>
</tr>
<tr>
<td>
<code>dataBase64</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DataBase64 is the base64-encoded string of the actual event&rsquo;s body posted to the sink. Default is empty.
Mutually exclusive with Data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.PingSourceStatus">PingSourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.PingSource">PingSource</a>)
</p>
<p>
<p>PingSourceStatus defines the observed state of PingSource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.SinkBindingSpec">SinkBindingSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.SinkBinding">SinkBinding</a>)
</p>
<p>
<p>SinkBindingSpec holds the desired state of the SinkBinding (from the client).</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>BindingSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#BindingSpec">
knative.dev/pkg/apis/duck/v1.BindingSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>BindingSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 BindingSpec, which currently provides:
* Subject - Subject references the resource(s) whose &ldquo;runtime contract&rdquo;
should be augmented by Binding implementations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1.SinkBindingStatus">SinkBindingStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1.SinkBinding">SinkBinding</a>)
</p>
<p>
<p>SinkBindingStatus communicates the observed state of the SinkBinding (from the controller).</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
<tr>
<td>
<code>oidcTokenSecretName</code><br/>
<em>
string
</em>
</td>
<td>
<p>OIDCTokenSecretName is the name of the secret containing the token for
this SinkBindings OIDC authentication</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="sources.knative.dev/v1alpha1">sources.knative.dev/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains API Schema definitions for the sources v1alpha1 API group.</p>
</p>
Resource Types:
<ul><li>
<a href="#sources.knative.dev/v1alpha1.IntegrationSource">IntegrationSource</a>
</li></ul>
<h3 id="sources.knative.dev/v1alpha1.IntegrationSource">IntegrationSource
</h3>
<p>
<p>IntegrationSource is the Schema for the Integrationsources API</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>IntegrationSource</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.IntegrationSourceSpec">
IntegrationSourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>aws</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.Aws">
Aws
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>timer</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.Timer">
Timer
</a>
</em>
</td>
<td>
<p>AWS source configuration</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.IntegrationSourceStatus">
IntegrationSourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1alpha1.Aws">Aws
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1alpha1.IntegrationSourceSpec">IntegrationSourceSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>s3</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSS3
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sqs</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSSQS
</em>
</td>
<td>
<p>S3 source configuration</p>
</td>
</tr>
<tr>
<td>
<code>ddbStreams</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.AWSDDBStreams
</em>
</td>
<td>
<p>SQS source configuration</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
knative.dev/eventing/pkg/apis/common/integration/v1alpha1.Auth
</em>
</td>
<td>
<p>DynamoDB Streams source configuration</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1alpha1.IntegrationSourceSpec">IntegrationSourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1alpha1.IntegrationSource">IntegrationSource</a>)
</p>
<p>
<p>IntegrationSourceSpec defines the desired state of IntegrationSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>aws</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.Aws">
Aws
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>timer</code><br/>
<em>
<a href="#sources.knative.dev/v1alpha1.Timer">
Timer
</a>
</em>
</td>
<td>
<p>AWS source configuration</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1alpha1.IntegrationSourceStatus">IntegrationSourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1alpha1.IntegrationSource">IntegrationSource</a>)
</p>
<p>
<p>IntegrationSourceStatus defines the observed state of IntegrationSource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1alpha1.Timer">Timer
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1alpha1.IntegrationSourceSpec">IntegrationSourceSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>period</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Interval (in milliseconds) between producing messages</p>
</td>
</tr>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message to generate</p>
</td>
</tr>
<tr>
<td>
<code>repeatCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>Content type of generated message</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="sources.knative.dev/v1beta2">sources.knative.dev/v1beta2</h2>
<p>
<p>Package v1beta2 contains API Schema definitions for the sources v1beta2 API group.</p>
</p>
Resource Types:
<ul><li>
<a href="#sources.knative.dev/v1beta2.PingSource">PingSource</a>
</li></ul>
<h3 id="sources.knative.dev/v1beta2.PingSource">PingSource
</h3>
<p>
<p>PingSource is the Schema for the PingSources API.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
sources.knative.dev/v1beta2
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>PingSource</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#sources.knative.dev/v1beta2.PingSourceSpec">
PingSourceSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>schedule</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schedule is the cron schedule. Defaults to <code>* * * * *</code>.</p>
</td>
</tr>
<tr>
<td>
<code>timezone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timezone modifies the actual time relative to the specified timezone.
Defaults to the system time zone.
More general information about time zones: <a href="https://www.iana.org/time-zones">https://www.iana.org/time-zones</a>
List of valid timezone values: <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">https://en.wikipedia.org/wiki/List_of_tz_database_time_zones</a></p>
</td>
</tr>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContentType is the media type of Data or DataBase64. Default is empty.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is data used as the body of the event posted to the sink. Default is empty.
Mutually exclusive with DataBase64.</p>
</td>
</tr>
<tr>
<td>
<code>dataBase64</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DataBase64 is the base64-encoded string of the actual event&rsquo;s body posted to the sink. Default is empty.
Mutually exclusive with Data.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#sources.knative.dev/v1beta2.PingSourceStatus">
PingSourceStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1beta2.PingSourceSpec">PingSourceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1beta2.PingSource">PingSource</a>)
</p>
<p>
<p>PingSourceSpec defines the desired state of the PingSource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceSpec</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceSpec">
knative.dev/pkg/apis/duck/v1.SourceSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceSpec</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceSpec, which currently provides:
* Sink - a reference to an object that will resolve to a domain name or
a URI directly to use as the sink.
* CloudEventOverrides - defines overrides to control the output format
and modifications of the event sent to the sink.</p>
</td>
</tr>
<tr>
<td>
<code>schedule</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Schedule is the cron schedule. Defaults to <code>* * * * *</code>.</p>
</td>
</tr>
<tr>
<td>
<code>timezone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timezone modifies the actual time relative to the specified timezone.
Defaults to the system time zone.
More general information about time zones: <a href="https://www.iana.org/time-zones">https://www.iana.org/time-zones</a>
List of valid timezone values: <a href="https://en.wikipedia.org/wiki/List_of_tz_database_time_zones">https://en.wikipedia.org/wiki/List_of_tz_database_time_zones</a></p>
</td>
</tr>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ContentType is the media type of Data or DataBase64. Default is empty.</p>
</td>
</tr>
<tr>
<td>
<code>data</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Data is data used as the body of the event posted to the sink. Default is empty.
Mutually exclusive with DataBase64.</p>
</td>
</tr>
<tr>
<td>
<code>dataBase64</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>DataBase64 is the base64-encoded string of the actual event&rsquo;s body posted to the sink. Default is empty.
Mutually exclusive with Data.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="sources.knative.dev/v1beta2.PingSourceStatus">PingSourceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#sources.knative.dev/v1beta2.PingSource">PingSource</a>)
</p>
<p>
<p>PingSourceStatus defines the observed state of PingSource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>SourceStatus</code><br/>
<em>
<a href="https://pkg.go.dev/knative.dev/pkg/apis/duck/v1#SourceStatus">
knative.dev/pkg/apis/duck/v1.SourceStatus
</a>
</em>
</td>
<td>
<p>
(Members of <code>SourceStatus</code> are embedded into this type.)
</p>
<p>inherits duck/v1 SourceStatus, which currently provides:
* ObservedGeneration - the &lsquo;Generation&rsquo; of the Service that was last
processed by the controller.
* Conditions - the latest available observations of a resource&rsquo;s current
state.
* SinkURI - the current active sink URI that has been configured for the
Source.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
.
</em></p>
