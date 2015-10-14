package com.comcast.cdn.traffic_control.traffic_router.core.util;

import com.comcast.cdn.traffic_control.traffic_router.core.loc.NetworkNodeException;

import java.net.Inet4Address;
import java.net.Inet6Address;
import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.Arrays;

public class CidrAddress implements Comparable<CidrAddress> {
    private final byte[] hostBytes;
    private final byte[] maskBytes;
    private final int netmaskLength;
    private final String cidrString;

    public static CidrAddress fromString(final String cidrString) throws NetworkNodeException {
        final String[] hostNetworkArray = cidrString.split("/");
        final String host = hostNetworkArray[0];

        InetAddress address;
        try {
            address = InetAddress.getByName(host);
        } catch (UnknownHostException ex) {
            throw new NetworkNodeException(ex);
        }

        if (hostNetworkArray.length == 1) {
            return new CidrAddress(address);
        }

        int netmaskLength;
        try {
            netmaskLength = Integer.parseInt(hostNetworkArray[1]);
        }
        catch (NumberFormatException e) {
            throw new NetworkNodeException(e);
        }

        return new CidrAddress(address, netmaskLength);
    }

    public CidrAddress(final InetAddress address) throws NetworkNodeException {
        this(address, address.getAddress().length * 8);
    }

    public CidrAddress(final InetAddress address, final int netmaskLength) throws NetworkNodeException {
        this.netmaskLength = netmaskLength;
        final byte[] addressBytes = address.getAddress();

        cidrString = address.toString() + "/" + netmaskLength;

        if (address instanceof Inet4Address && (netmaskLength > 32 || netmaskLength < 0)) {
            throw new NetworkNodeException("Rejecting IPv4 subnet with invalid netmask: " + cidrString);
        } else if (address instanceof Inet6Address && (netmaskLength > 128 || netmaskLength < 0)) {
            throw new NetworkNodeException("Rejecting IPv6 subnet with invalid netmask: " + cidrString);
        }

        hostBytes = addressBytes;
        maskBytes = new byte[addressBytes.length];

        for (int i = 0; i < netmaskLength; i++) {
            maskBytes[i/8] |= 1<<(7-(i%8));
        }
    }

    public byte[] getHostBytes() {
        return hostBytes;
    }

    public byte[] getMaskBytes() {
        return maskBytes;
    }

    public int getNetmaskLength() {
        return netmaskLength;
    }

    public boolean includesAddress(final CidrAddress other) {
        return compareTo(other) == 0;
    }

    public boolean isIpV6() {
        return getHostBytes().length > 4;
    }

    @Override
	public int compareTo(final CidrAddress other) {
		byte[] mask = this.maskBytes;
		int len = netmaskLength;

		if (netmaskLength > other.netmaskLength) {
			mask = other.maskBytes;
			len = other.netmaskLength;
		}

		final int numNetmaskBytes = (int) Math.ceil((double) len / 8);

        for (int i = 0; i < numNetmaskBytes; i++) {
            final int diff = (hostBytes[i] & mask[i]) - (other.hostBytes[i] & mask[i]);
			if (diff != 0) {
                return diff;
            }
		}

        return 0;
    }

    @Override
    @SuppressWarnings({"PMD.NPathComplexity", "PMD.IfStmtsMustUseBraces"})
    public boolean equals(final Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;

        final CidrAddress that = (CidrAddress) o;

        if (netmaskLength != that.netmaskLength) return false;
        if (!Arrays.equals(hostBytes, that.hostBytes)) return false;
        if (!Arrays.equals(maskBytes, that.maskBytes)) return false;
        return true;
    }

    @Override
    public int hashCode() {
        int result = hostBytes != null ? Arrays.hashCode(hostBytes) : 0;
        result = 31 * result + (maskBytes != null ? Arrays.hashCode(maskBytes) : 0);
        result = 31 * result + netmaskLength;
        return result;
    }

    @Override
    public String toString() {
        return "CidrAddress{" + cidrString + "}";
    }
}