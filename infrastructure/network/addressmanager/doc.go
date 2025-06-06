/*
Package addressmanager implements concurrency safe Waglayla address manager.

# Address Manager Overview

In order maintain the peer-to-peer Waglayla network, there needs to be a source
of addresses to connect to as nodes come and go. The Waglayla protocol provides
the getaddr and addr messages to allow peers to communicate known addresses with
each other. However, there needs to a mechanism to store those results and
select peers from them. It is also important to note that remote peers can't
be trusted to send valid peers nor attempt to provide you with only peers they
control with malicious intent.

With that in mind, this package provides a concurrency safe address manager for
caching and selecting peers in a non-deterministic manner. The general idea is
the caller adds addresses to the address manager and notifies it when addresses
are connected, known good, and attempted. The caller also requests addresses as
it needs them.

The address manager internally segregates the addresses into groups and
non-deterministically selects groups in a cryptographically random manner. This
reduce the chances multiple addresses from the same nets are selected which
generally helps provide greater peer diversity, and perhaps more importantly,
drastically reduces the chances an attacker is able to coerce your peer into
only connecting to nodes they control.

The address manager also understands routability and tries hard to only return
routable addresses. In addition, it uses the information provided by the caller
about connected, known good, and attempted addresses to periodically purge
peers which no longer appear to be good peers as well as bias the selection
toward known good peers. The general idea is to make a best effort at only
providing usable addresses.
*/
package addressmanager
